package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"todo-server/config"
	"todo-server/internal/database"
	"todo-server/internal/handlers"
	"todo-server/internal/middleware"
	"todo-server/internal/repository"
	"todo-server/internal/routes"
)

//go:embed migrations/*.sql
var migrationFiles embed.FS

// Simple response structure (like a DTO in TypeScript)
type HealthResponse struct {
	Status    string    `json:"status"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
	Version   string    `json:"version"`
}

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Connect to DB
	db, err := database.NewPostgresDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err)
	}
	defer db.Close()

	// Run Migrations
	if err := database.RunMigrations(db, migrationFiles); err != nil {
		log.Fatalf("Failed to run database migrations: %v", err)
	}

	// Create repositories: Postgres + Cached wrapper
	postgresRepo := repository.NewPostgresRepository(db)
	cachedRepo := repository.NewCachedRepository(postgresRepo)

	// Create handler with cached repository
	todoHandler := handlers.NewTodoHandler(cachedRepo)

	// Setup routes
	mux := routes.SetupRoutesWithHandler(todoHandler)

	// Apply middleware
	handler := middleware.LoggingMiddleware(
		middleware.CORSMiddleware(mux),
	)

	port := os.Getenv("PORT")

	if port == "" {
		port = ":9000" // fallback for local development
	}

	server := &http.Server{
		Addr:         port,
		Handler:      handler,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	fmt.Printf("🚀 Go Todo API started on http://localhost%s\n", port)
	fmt.Printf("Environment: %s | DB Port: %d | Cache: Enabled\n", cfg.Environment, cfg.DBPort)
	fmt.Println("Using PostgreSQL + In-Memory Cache")

	log.Printf("Server starting on port %s", port)
	log.Fatal(server.ListenAndServe())
}

// ---------- OLD CODE FOR LEARNING --------------
// func main() {
// 	// Create a new ServeMux (this is our router, similar to Express Router)
// 	mux := routes.SetupRoutes()

// 	handler := middleware.LoggingMiddleware(middleware.CORSMiddleware(mux))

// 	// Start the server on port 9000
// 	port := ":9000"
// 	server := &http.Server{
// 		Addr:         port,
// 		Handler:      handler,
// 		ReadTimeout:  10 * time.Second,
// 		WriteTimeout: 10 * time.Second,
// 		IdleTimeout:  6 * time.Second,
// 	}

// 	fmt.Printf("🚀 Go Todo API started on http://localhost%s\n", port)
// 	fmt.Println("Available routes:")
// 	fmt.Println("   GET  /")
// 	fmt.Println("   GET  /health")
// 	fmt.Println("   GET  /todos")
// 	fmt.Println("   POST /todos")
// 	fmt.Println("   GET  /todos/{id}")
// 	fmt.Println("   PUT  /todos/{id}")
// 	fmt.Println("   PATCH  /todos/{id}")
// 	fmt.Println("   DELETE /todos/{id}")
// 	fmt.Println("   DELETE /todos")

// 	log.Fatalln(server.ListenAndServe())
// }
