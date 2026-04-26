package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	ServerPort    string
	Environment   string
	AllowedOrigin string

	// Database
	DBHost     string
	DBPort     int
	DBName     string
	DBUser     string
	DBPassword string
}

func LoadConfig() (*Config, error) {
	dbPortStr := getEnv("DB_PORT", "5434")

	dbPort, error := strconv.Atoi(dbPortStr)

	if error != nil {
		return nil, fmt.Errorf("Invalid DB_PORT: %w: ", error)
	}

	cfg := &Config{
		ServerPort:    getEnv("SERVER_PORT", ":9000"),
		Environment:   getEnv("ENVIRONMENT", "development"),
		AllowedOrigin: getEnv("ALLOWED_ORIGIN", "http://localhost:5173"),

		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     dbPort,
		DBName:     getEnv("DB_NAME", "todo_db"),
		DBUser:     getEnv("DB_USER", "todo_user"),
		DBPassword: getEnv("DB_PASSWORD", "todo_password"),
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
