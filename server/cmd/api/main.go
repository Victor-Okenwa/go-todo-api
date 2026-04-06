package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Simple response structure (like a DTO in TypeScript)
type HealthResponse struct {
	Status    string    `json:"status"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
	Version   string    `json:"version"`
}

func main() {
	// Create a new ServeMux (this is our router, similar to Express Router)
	mux := http.NewServeMux()

	// Basic route - Welcome message (like GET /api in Express)

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Welcome to the Go API!",
			"info":    "Try visiting /health for a health check.",
		})
	})

	// Basic route - Health check (like GET /api/health in Express)
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		response := HealthResponse{
			Status:    "ok",
			Message:   "API is healthy",
			Timestamp: time.Now(),
			Version:   "0.1.0",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "Failed to Encode Response", http.StatusInternalServerError)
			return
		}
	})

	mux.HandleFunc("GET /todos", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]any{
			"message": "This is where your todos would be returned.",
			"count":   0,
			"todos":   []string{},
		})
	})

	// Start the server on port 9000
	port := ":9000"
	server := &http.Server{
		Addr:         port,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  6 * time.Second,
	}

	fmt.Printf("🚀 Go Todo API server starting on http://localhost%s\n", port)
	fmt.Println("Available routes:")
	fmt.Println("   GET /          → Welcome message")
	fmt.Println("   GET /health    → Health check")
	fmt.Println("   GET /todos     → Todo list (placeholder)")

	log.Fatalln(server.ListenAndServe())
}
