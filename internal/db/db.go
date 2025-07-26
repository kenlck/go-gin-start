// Package db provides database connection logic using pgxpool.
package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func InitDB(databaseURL string) error {
	config, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		return err
	}
	DB, err = pgxpool.NewWithConfig(context.Background(), config)
	return err
}
