package routes

import (
	"encoding/json"
	"net/http"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

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

	// todos route (placeholder)
	mux.HandleFunc("GET /todos", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]any{
			"message": "This is where your todos would be returned.",
			"count":   0,
			"todos":   []string{},
		})
	})

	// Apply middleware (CORS + Logging)
	// Note: Middleware order matters - Logging outermost, CORS inner
	return mux
}
