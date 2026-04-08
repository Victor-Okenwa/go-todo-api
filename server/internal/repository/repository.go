package repository

import "todo-server/internal/models"

type CheckedState struct {
	Completed bool `json:"completed"`
}

type TodoRepository interface {
	GetAll() ([]models.Todo, error)
	GetByID(id int) (models.Todo, error)
	Create(todo models.Todo) (models.Todo, error)
	Update(id int, updatedTodo models.Todo) (models.Todo, error)
	UpdateCompleted(id int, checkedState CheckedState) (models.Todo, error)
	Delete(id int) error
	DeleteAll() error
}
