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
