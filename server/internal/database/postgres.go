package database

import (
	"context"
	"fmt"
	"log"
	"time"
	"todo-server/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	Pool *pgxpool.Pool
}

// NewPostgresDB creates a new PostgreSQL connection pool
func NewPostgresDB(cfg *config.Config) (*DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	// Connection pool config
	poolConfig, error := pgxpool.ParseConfig(dsn)

	if error != nil {
		return nil, fmt.Errorf("Failed to parse DB config: %w", error)
	}

	// Reasonable defaults for a small todo app
	poolConfig.MaxConns = 10
	poolConfig.MinConns = 2
	poolConfig.MaxConnLifetime = 30 * time.Minute // 30 minutes

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pool, error := pgxpool.NewWithConfig(ctx, poolConfig)

	if error != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", error)
	}

	// Test the connection
	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("✅ Successfully connected to PostgreSQL")
	return &DB{Pool: pool}, nil
}

func (db *DB) Close() {
	if db.Pool != nil {
		db.Pool.Close()
		log.Println("🔌 Database connection pool closed")
	}
}
