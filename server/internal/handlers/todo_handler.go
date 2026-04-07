package handlers

import (
	"encoding/json"
	"net/http"
	"todo-server/internal/models"
	"todo-server/internal/repository"
)

type TodoHandler struct {
	repo repository.TodoRepository
}

// NewTodoHandler creates a new handler with the repository
func NewTodoHandler(repo repository.TodoRepository) *TodoHandler {
	return &TodoHandler{repo: repo}
}

// GetAllTodos returns all todos
func (h *TodoHandler) GetAllTodos(w http.ResponseWriter, r *http.Request) {
	todo, err := h.repo.GetAll()
	if err != nil {
		http.Error(w, "Failed to get todos", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}



// CreateTodo creates a new todo
func (h *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Use NewTodo to ensure proper defaults
	newTodo := models.NewTodo(todo.Title, todo.Description)

	created, err := h.repo.Create(newTodo)

	if err != nil {
		http.Error(w, "Failed to create todo", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
}
