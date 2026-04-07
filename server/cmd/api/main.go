package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"todo-server/internal/middleware"
	"todo-server/internal/routes"
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
	mux := routes.SetupRoutes()

	handler := middleware.LoggingMiddleware(middleware.CORSMiddleware(mux))

	// Start the server on port 9000
	port := ":9000"
	server := &http.Server{
		Addr:         port,
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  6 * time.Second,
	}

	fmt.Printf("🚀 Go Todo API started on http://localhost%s\n", port)
	fmt.Println("Available routes:")
	fmt.Println("   GET  /")
	fmt.Println("   GET  /health")
	fmt.Println("   GET  /todos")
	fmt.Println("   POST /todos")
	fmt.Println("   GET  /todos/{id}")
	fmt.Println("   PUT  /todos/{id}")
	fmt.Println("   DELETE /todos/{id}")
	fmt.Println("   DELETE /todos")

	log.Fatalln(server.ListenAndServe())
}
