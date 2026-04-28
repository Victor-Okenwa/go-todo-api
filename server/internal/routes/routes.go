package routes

import (
	"net/http"
	"todo-server/internal/handlers"
)

func SetupRoutesWithHandler(todoHandler *handlers.TodoHandler) *http.ServeMux {
	mux := http.NewServeMux()

	// Health check
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"ok","message":"Todo API is running with Postgres + Cache!"}`))
	})

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message":"Welcome to Go Todo API (Postgres + Cache)"}`))
	})

	// Todo Routes
	mux.HandleFunc("GET /todos", todoHandler.GetAllTodos)
	mux.HandleFunc("POST /todos", todoHandler.CreateTodo)
	mux.HandleFunc("GET /todos/{id}", todoHandler.GetByID)
	mux.HandleFunc("PUT /todos/{id}", todoHandler.UpdateTodo)
	mux.HandleFunc("PATCH /todos/{id}", todoHandler.UpdateCompleted)
	mux.HandleFunc("DELETE /todos/{id}", todoHandler.Delete)
	mux.HandleFunc("DELETE /todos", todoHandler.DeleteAll)

	return mux
}
