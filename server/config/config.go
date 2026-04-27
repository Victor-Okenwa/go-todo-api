package config

import (
	"os"
	"strconv"
)

type Config struct {
	ServerPort    string
	Environment   string
	AllowedOrigin string

	// Database - Support two modes
	UseSupabase        bool
	DBConnectionString string // Full postgres:// URL (for Supabase or local)

	// Database
	DBHost     string
	DBPort     int
	DBName     string
	DBUser     string
	DBPassword string
}

func LoadConfig() (*Config, error) {
	env := getEnv("ENVIRONMENT", "development")

	cfg := &Config{
		ServerPort:    getEnv("SERVER_PORT", ":9000"),
		Environment:   env,
		AllowedOrigin: getEnv("ALLOWED_ORIGIN", "http://localhost:5173"),
	}

	// Check if we should use Supabase connection string (preferred for prod)
	if dsn := os.Getenv("DB_CONNECTION_STRING"); dsn != "" {
		cfg.UseSupabase = true
		cfg.DBConnectionString = dsn
	} else {
		// Fallback to local Docker Postgres
		cfg.UseSupabase = false
		dbPort, _ := strconv.Atoi(getEnv("DB_PORT", "5434"))
		cfg.DBHost = getEnv("DB_HOST", "localhost")
		cfg.DBPort = dbPort
		cfg.DBName = getEnv("DB_NAME", "todo_db")
		cfg.DBUser = getEnv("DB_USER", "todo_user")
		cfg.DBPassword = getEnv("DB_PASSWORD", "todo_password")
	}

	return cfg, nil
}

// Helper to get env with default
func getEnv(key, defaultVal string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return defaultVal
}
