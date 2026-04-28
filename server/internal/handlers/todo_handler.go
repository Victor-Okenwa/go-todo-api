package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
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

// Get By ID returns a single todo by ID
func (h *TodoHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Id is required", http.StatusBadRequest)
		return
	}

	todo, err := h.repo.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusFound)
	json.NewEncoder(w).Encode(todo)
}

// CreateTodo creates a new todo
func (h *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Basic validation
	if strings.TrimSpace(todo.Title) == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}
	// Use NewTodo to ensure proper defaults
	newTodo := models.NewTodo(todo.Title, todo.Description)
	created, err := h.repo.Create(newTodo)

	if err != nil {
		log.Printf("ERROR creating todo: %v", err)
		http.Error(w, "Failed to create todo", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
}

func (h *TodoHandler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadGateway)
		return
	}

	var update models.Todo

	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	fmt.Printf("Received update for ID %d: %+v\n", id, update)

	result, err := h.repo.Update(id, update)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(result)
}

func (h *TodoHandler) UpdateCompleted(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadGateway)
		return
	}

	var completed repository.CheckedState

	if err := json.NewDecoder(r.Body).Decode(&completed); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	result, err := h.repo.UpdateCompleted(id, completed)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(result)
}

// DELETE BY ID
func (h *TodoHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Id", http.StatusBadRequest)
		return
	}

	if err := h.repo.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DeleteAllTodos deletes all todos
func (h *TodoHandler) DeleteAll(w http.ResponseWriter, r *http.Request) {
	if err := h.repo.DeleteAll(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
