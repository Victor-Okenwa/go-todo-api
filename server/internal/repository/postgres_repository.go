package repository

import (
	"context"
	"errors"
	"todo-server/internal/database"
	"todo-server/internal/models"

	"github.com/jackc/pgx/v5"
)

type postgresRepository struct {
	db *database.DB
}

func NewPostgresRepository(db *database.DB) TodoRepository {
	return &postgresRepository{db: db}
}

func (r *postgresRepository) GetAll() ([]models.Todo, error) {
	ctx := context.Background()
	query := `SELECT id, title, description, completed, created_at, updated_at 
	          FROM todos ORDER BY created_at DESC`
	rows, err := r.db.Pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var todos []models.Todo

	for rows.Next() {
		var t models.Todo
		err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Completed, &t.CreatedAt, &t.UpdatedAt)
		if err != nil {
			return nil, err
		}
		todos = append(todos, t)
	}
	return todos, nil
}

func (r *postgresRepository) GetByID(id int) (models.Todo, error) {
	ctx := context.Background()
	query := `SELECT * FROM todos WHERE id = $1`

	var t models.Todo

	err := r.db.Pool.QueryRow(ctx, query, id).Scan(&t.ID, &t.Title, &t.Description, &t.Completed, &t.CreatedAt, &t.UpdatedAt)
	if err != nil {
		// checking if error is coming from no ID found
		if errors.Is(err, pgx.ErrNoRows) {
			return models.Todo{}, errors.New("todo not found")
		}

		return models.Todo{}, err
	}
	return t, nil
}

func (r *postgresRepository) Create(todo models.Todo) (models.Todo, error) {
	ctx := context.Background()
	query := `INSERT INTO todos (title, description, completed) 
	          VALUES ($1, $2, $3) RETURNING id, created_at, updated_at`

	err := r.db.Pool.QueryRow(ctx, query, todo.Title, todo.Description, todo.Completed).Scan(&todo.ID, &todo.Title, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt)
	if err != nil {
		return models.Todo{}, err
	}

	return todo, nil
}

func (r *postgresRepository) Update(id int, updatedTodo models.Todo) (models.Todo, error) {
	ctx := context.Background()
	query := `UPDATE todos SET title = $1, description = $2, completed = $3, updated_at = NOW() 
	          WHERE id = $4 RETURNING id, title, description, completed, created_at, updated_at`

	var t models.Todo

	err := r.db.Pool.QueryRow(ctx, query, updatedTodo.Title, updatedTodo.Description, updatedTodo.Completed).Scan(&t.ID, &t.Title, &t.Description, &t.Completed, &t.CreatedAt, &t.UpdatedAt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.Todo{}, errors.New("todo not found")
		}
		return models.Todo{}, err
	}

	return t, nil
}

func (r *postgresRepository) UpdateCompleted(id int, state CheckedState) (models.Todo, error) {
	ctx := context.Background()
	query := `UPDATE todos SET completed = $1, updated_at = NOW() 
		  WHERE id = $2 RETURNING id, title, description, completed, created_at, updated_at`
	var t models.Todo
	err := r.db.Pool.QueryRow(ctx, query, state.Completed, id).Scan(&t.ID, &t.Title, &t.Description, &t.Completed, &t.CreatedAt, &t.UpdatedAt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.Todo{}, errors.New("todo not found")
		}
		return models.Todo{}, err
	}

	return t, nil
}

func (r *postgresRepository) Delete(id int) error {
	ctx := context.Background()
	query := `DELETE FROM todos WHERE id = $1`

	result, err := r.db.Pool.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errors.New("todo not found")
	}
	return nil
}

func (r *postgresRepository) DeleteAll() error {
	ctx := context.Background()
	query := `DELETE FROM todos`
	_, err := r.db.Pool.Exec(ctx, query)

	if err != nil {
		return err
	}
	return nil
}
