// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: properties.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
	null "github.com/guregu/null/v5"
	"github.com/lib/pq"
)

const createProperties = `-- name: CreateProperties :one
INSERT INTO properties (id, created_at, updated_at, type, address, city, bedrooms, bathrooms, price, square_meters, description, long_description, location, is_available, realtors, images)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)
RETURNING id, created_at, updated_at, type, address, city, bedrooms, bathrooms, price, square_meters, description, long_description, location, is_available, realtors, images
`

type CreatePropertiesParams struct {
	ID              uuid.UUID
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Type            string
	Address         string
	City            string
	Bedrooms        null.Int32
	Bathrooms       null.Int32
	Price           int32
	SquareMeters    int32
	Description     string
	LongDescription string
	Location        null.String
	IsAvailable     bool
	Realtors        []string
	Images          []string
}

func (q *Queries) CreateProperties(ctx context.Context, arg CreatePropertiesParams) (Property, error) {
	row := q.db.QueryRowContext(ctx, createProperties,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Type,
		arg.Address,
		arg.City,
		arg.Bedrooms,
		arg.Bathrooms,
		arg.Price,
		arg.SquareMeters,
		arg.Description,
		arg.LongDescription,
		arg.Location,
		arg.IsAvailable,
		pq.Array(arg.Realtors),
		pq.Array(arg.Images),
	)
	var i Property
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Type,
		&i.Address,
		&i.City,
		&i.Bedrooms,
		&i.Bathrooms,
		&i.Price,
		&i.SquareMeters,
		&i.Description,
		&i.LongDescription,
		&i.Location,
		&i.IsAvailable,
		pq.Array(&i.Realtors),
		pq.Array(&i.Images),
	)
	return i, err
}
