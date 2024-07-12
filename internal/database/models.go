// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package database

import (
	"time"

	"github.com/google/uuid"
	null "github.com/guregu/null/v5"
)

type Property struct {
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

type User struct {
	ID           uuid.UUID
	CreatedAt    time.Time
	UpdatedAt    time.Time
	FirstName    string
	LastName     string
	EmailAddress string
	UserName     string
	IsAdmin      bool
}