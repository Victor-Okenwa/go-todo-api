package routes

import (
	"encoding/json"
	"net/http"
	"todo-server/internal/handlers"
	"todo-server/internal/repository"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	repo := repository.NewMemoryRepository()
	todoHandlers := handlers.NewTodoHandler(repo)

	// Root Route
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Welcome to the Go API!",
			"info":    "Try visiting /health for a health check.",
		})
	})

	// health check route
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		response := map[string]any{
			"status":    "ok",
			"message":   "API is healthy",
			"timestamp": "2024-06-01T12:00:00Z",
			"version":   "0.1.0",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	})

	// Todo Routes
	mux.HandleFunc("GET /todos", todoHandlers.GetAllTodos)
	mux.HandleFunc("GET /todos/{id}", todoHandlers.GetByID)
	mux.HandleFunc("POST /todos", todoHandlers.CreateTodo)
	mux.HandleFunc("PUT /todos/{id}", todoHandlers.UpdateTodo)
	mux.HandleFunc("DELETE /todos/{id}", todoHandlers.Delete)
	mux.HandleFunc("DELETE /todos", todoHandlers.DeleteAll)

	return mux
}
