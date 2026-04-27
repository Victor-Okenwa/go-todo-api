package database

import (
	"context"
	"embed"
	"fmt"
	"log"
)

func RunMigrations(db *DB, migrationFiles embed.FS) error {
	// Read the migration files from the embedded filesystem
	sqlBytes, err := migrationFiles.ReadFile("migrations/001_create_todos_table.sql")
	if err != nil {
		return fmt.Errorf("failed to read embedded migration file: %w", err)
	}

	ctx := context.Background()
	_, err = db.Pool.Exec(ctx, string(sqlBytes))

	if err != nil {
		return fmt.Errorf("failed to execute migration: %w", err)
	}

	log.Println("✅ Database migration completed successfully (table 'todos' is ready)")
	return nil
}
