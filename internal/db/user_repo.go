// Package db provides user repository logic using Squirrel and pgxpool.
package db

import (
	"context"

	"github.com/Masterminds/squirrel"

	"go-gin-start/internal/model"
)

// UserRepo handles user CRUD operations.
type UserRepo struct{}

// FindByUsername returns a user by username.
func (r *UserRepo) FindByUsername(ctx context.Context, username string) (*model.User, error) {
	sql, args, err := squirrel.Select("id", "username", "password_hash", "created_at").
		From("users").
		Where(squirrel.Eq{"username": username}).
		ToSql()
	if err != nil {
		return nil, err
	}
	row := DB.QueryRow(ctx, sql, args...)
	var u model.User
	err = row.Scan(&u.ID, &u.Username, &u.PasswordHash, &u.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

// CreateUser inserts a new user.
func (r *UserRepo) CreateUser(ctx context.Context, username, passwordHash string) error {
	sql, args, err := squirrel.Insert("users").
		Columns("username", "password_hash").
		Values(username, passwordHash).
		ToSql()
	if err != nil {
		return err
	}
	_, err = DB.Exec(ctx, sql, args...)
	return err
}
