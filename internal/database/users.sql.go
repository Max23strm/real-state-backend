// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: users.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createUsers = `-- name: CreateUsers :one
INSERT INTO users (id, created_at, updated_at, first_name, last_name, email_address, user_name, is_admin)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id, created_at, updated_at, first_name, last_name, email_address, user_name, is_admin
`

type CreateUsersParams struct {
	ID           uuid.UUID
	CreatedAt    time.Time
	UpdatedAt    time.Time
	FirstName    string
	LastName     string
	EmailAddress string
	UserName     string
	IsAdmin      bool
}

func (q *Queries) CreateUsers(ctx context.Context, arg CreateUsersParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUsers,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.FirstName,
		arg.LastName,
		arg.EmailAddress,
		arg.UserName,
		arg.IsAdmin,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.FirstName,
		&i.LastName,
		&i.EmailAddress,
		&i.UserName,
		&i.IsAdmin,
	)
	return i, err
}
