package models

import (
	"time"
)

type Todo struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// Creates new todo
func NewTodo(title, description string) Todo {
	now := time.Now()

	return Todo{
		Title:       title,
		Description: description,
		Completed:   false,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

// Update updates the todo fields and refreshes UpdatedAt
func (t *Todo) Update(title, description string, completed bool) {
	if t.Title != "" {
		t.Title = title
	}
	if t.Description != "" {
		t.Description = description
	}

	t.Completed = completed
	t.UpdatedAt = time.Now()
}

// Update checked (completed)
func (t *Todo) UpdateCompleted(completed bool) {
	t.Completed = completed
	t.UpdatedAt = time.Now()
}
