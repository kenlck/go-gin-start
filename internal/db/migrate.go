// Package db provides migration logic using golang-migrate.
package db

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
)

// AutoMigrate applies all migrations in the migrations/ directory.
func AutoMigrate(databaseURL string) error {
	m, err := migrate.New(
		"file://migrations",
		databaseURL,
	)
	if err != nil {
		return fmt.Errorf("failed to initialize migrate: %w", err)
	}
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("migration failed: %w", err)
	}
	m.Close()
	return nil
}
