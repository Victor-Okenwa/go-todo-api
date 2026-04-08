package repository

import (
	"errors"
	"sync"
	"todo-server/internal/models"
)

// memoryRepository is an in-memory implementation using a slice + mutex for safety
type memoryRepository struct {
	todos  []models.Todo
	nextID int
	mu     sync.Mutex // protects concurrent access (important because of goroutines)
}

func NewMemoryRepository() TodoRepository {
	return &memoryRepository{
		todos:  make([]models.Todo, 0),
		nextID: 1,
	}
}

func (r *memoryRepository) GetAll() ([]models.Todo, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	//  return a copy of the slice to prevent external modification
	result := make([]models.Todo, len(r.todos))
	copy(result, r.todos)
	return result, nil
}

func (r *memoryRepository) GetByID(id int) (models.Todo, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, todo := range r.todos {
		if todo.ID == id {
			return todo, nil
		}
	}

	return models.Todo{}, errors.New("New Todo Found")
}

func (r *memoryRepository) Create(todo models.Todo) (models.Todo, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	todo.ID = r.nextID
	r.nextID++
	r.todos = append(r.todos, todo)

	return todo, nil
}

func (r *memoryRepository) Update(id int, updatedTodo models.Todo) (models.Todo, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i, todo := range r.todos {
		if todo.ID == id {

			if updatedTodo.Title == "" {
				address := &updatedTodo.Title
				*address = todo.Title
			}
			r.todos[i].Update(updatedTodo.Title, updatedTodo.Description, updatedTodo.Completed)
			return r.todos[i], nil
		}
	}

	return models.Todo{}, errors.New("Todo Not Found")
}

func (r *memoryRepository) UpdateCompleted(id int, state CheckedState) (models.Todo, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i, todo := range r.todos {
		if todo.ID == id {
			r.todos[i].UpdateCompleted(state.Completed)
			return r.todos[i], nil
		}
	}

	return models.Todo{}, errors.New("Todo Not Found")
}

func (r *memoryRepository) Delete(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i, todo := range r.todos {
		if todo.ID == id {
			// Remove the todo from the slice
			r.todos = append(r.todos[:i], r.todos[i+1:]...)
		}
	}
	return errors.New("Todo Not Found")
}

func (r *memoryRepository) DeleteAll() error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.todos = make([]models.Todo, 0)
	r.nextID = 1
	return nil
}
