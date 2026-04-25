package database

import (
	"context"
	"fmt"
	"log"
	"os"
)

func RunMigrations(db *DB) error {
	migrationSQL, error := os.ReadFile("../../migrations/001_create_todos_table.sql")

	if error != nil {
		return fmt.Errorf("failed to read migration file: %w", error)
	}

	ctx := context.Background()

	_, error = db.Pool.Exec(ctx, string(migrationSQL))
	if error != nil {
		return fmt.Errorf("failed to run migration: %w", error)
	}

	log.Println("✅ Database migration completed successfully")
	return nil
}
