package main

import (
	"time"

	"github.com/Max23strm/real-state-backend/internal/database"
	"github.com/google/uuid"
	"github.com/guregu/null/v5"
)

type User struct {
	ID           uuid.UUID `json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	FirstName    string    `field:"first_name"`
	LastName     string    `field:"lastName"`
	EmailAddress string    `field:"email_address"`
	UserName     string    `field:"user_name"`
	IsAdmin      bool      `field:"is_admin"`
}

func databaseUserToUser(dbUser database.User) User {

	return User{
		ID:           dbUser.ID,
		CreatedAt:    dbUser.CreatedAt,
		UpdatedAt:    dbUser.UpdatedAt,
		FirstName:    dbUser.FirstName,
		LastName:     dbUser.LastName,
		EmailAddress: dbUser.EmailAddress,
		UserName:     dbUser.UserName,
		IsAdmin:      dbUser.IsAdmin,
	}
}

type Property struct {
	ID              uuid.UUID   `json:"id"`
	CreatedAt       time.Time   `json:"created_at"`
	UpdatedAt       time.Time   `json:"updated_at"`
	Type            string      `field:"type"`
	Address         string      `field:"address"`
	City            string      `field:"city"`
	Bedrooms        null.Int32  `field:"bedrooms"`
	Bathrooms       null.Int32  `field:"bathrooms"`
	Price           int32       `field:"price"`
	SquareMeters    int32       `field:"square_meters"`
	Description     string      `field:"description"`
	LongDescription string      `field:"long_description"`
	Location        null.String `field:"location"`
	IsAvailable     bool        `field:"is_available"`
	Realtors        []string    `field:"realtors"`
	Images          []string    `field:"images"`
}

func databasePropertyToProperty(dbProperty database.Property) Property {

	return Property{
		ID:              dbProperty.ID,
		CreatedAt:       dbProperty.CreatedAt,
		UpdatedAt:       dbProperty.UpdatedAt,
		Type:            dbProperty.Type,
		Address:         dbProperty.Address,
		City:            dbProperty.City,
		Bedrooms:        dbProperty.Bedrooms,
		Bathrooms:       dbProperty.Bathrooms,
		Price:           dbProperty.Price,
		SquareMeters:    dbProperty.SquareMeters,
		Description:     dbProperty.Description,
		LongDescription: dbProperty.LongDescription,
		Location:        dbProperty.Location,
		IsAvailable:     dbProperty.IsAvailable,
		Realtors:        dbProperty.Realtors,
		Images:          dbProperty.Images,
	}
}
