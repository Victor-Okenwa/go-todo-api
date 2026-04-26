package repository

import (
	"sync"
	"todo-server/internal/models"
)

// CachedRepository wraps Postgres repository and adds in-memory caching
type CachedRepository struct {
	postgres TodoRepository
	cache    map[int]models.Todo
	mu       sync.RWMutex
}

func NewCachedRepository(postgresRepo TodoRepository) TodoRepository {
	return &CachedRepository{
		postgres: postgresRepo,
		cache:    make(map[int]models.Todo),
	}
}

// GetAll - For simplicity, we cache the entire list (common for small todo apps)
func (c *CachedRepository) GetAll() ([]models.Todo, error) {
	c.mu.RLock()
	if len(c.cache) > 0 {
		c.mu.RUnlock()
		// Return copy of cached data
		todos := make([]models.Todo, 0, len(c.cache))
		for _, t := range c.cache {
			todos = append(todos, t)
		}
		return todos, nil
	}
	c.mu.RUnlock()

	// Cache miss → fetch from Postgres
	todos, err := c.postgres.GetAll()
	if err != nil {
		return nil, err
	}

	// Populate cache
	c.mu.Lock()
	c.cache = make(map[int]models.Todo, len(todos))
	for _, t := range todos {
		c.cache[t.ID] = t
	}
	c.mu.Unlock()

	return todos, nil
}

func (c *CachedRepository) GetByID(id int) (models.Todo, error) {
	// Check cache first (fast path)
	c.mu.RLock()
	if todo, exists := c.cache[id]; exists {
		c.mu.RUnlock()
		return todo, nil
	}
	c.mu.RUnlock()

	// Cache miss → go to database
	todo, err := c.postgres.GetByID(id)
	if err != nil {
		return models.Todo{}, err
	}

	// Store in cache
	c.mu.Lock()
	c.cache[id] = todo
	c.mu.Unlock()

	return todo, nil
}

// Write operations - always write to Postgres then invalidate cache
func (c *CachedRepository) Create(todo models.Todo) (models.Todo, error) {
	created, err := c.postgres.Create(todo)
	if err != nil {
		return models.Todo{}, err
	}

	// Invalidate cache (simple strategy: clear entire cache)
	c.mu.Lock()
	c.cache = make(map[int]models.Todo)
	c.mu.Unlock()

	return created, nil
}

func (c *CachedRepository) Update(id int, updated models.Todo) (models.Todo, error) {
	result, err := c.postgres.Update(id, updated)
	if err != nil {
		return models.Todo{}, err
	}

	c.mu.Lock()
	delete(c.cache, id)  // Remove old version
	c.cache[id] = result // Store new version
	c.mu.Unlock()

	return result, nil
}

func (c *CachedRepository) UpdateCompleted(id int, state CheckedState) (models.Todo, error) {
	result, err := c.postgres.UpdateCompleted(id, state)
	if err != nil {
		return models.Todo{}, err
	}

	c.mu.Lock()
	delete(c.cache, id)
	c.cache[id] = result // Store new version
	c.mu.Unlock()

	return result, nil
}
func (c *CachedRepository) Delete(id int) error {
	err := c.postgres.Delete(id)
	if err != nil {
		return err
	}

	c.mu.Lock()
	delete(c.cache, id)
	c.mu.Unlock()

	return nil
}

func (c *CachedRepository) DeleteAll() error {
	err := c.postgres.DeleteAll()
	if err != nil {
		return err
	}

	c.mu.Lock()
	c.cache = make(map[int]models.Todo)
	c.mu.Unlock()

	return nil
}
