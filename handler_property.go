package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Max23strm/real-state-backend/internal/database"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/guregu/null/v5"
)

func (apiCgf *apiConfig) handlerCreateProperty(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		Type            string      `field:"type" validate:"required"`
		Address         string      `field:"address" validate:"required,min=8"`
		City            string      `field:"city" validate:"required"`
		Bedrooms        null.Int32  `field:"bedrooms"`
		Bathrooms       null.Int32  `field:"bathrooms"`
		Price           int32       `field:"price" validate:"required,number"`
		SquareMeters    int32       `field:"square_meters" validate:"required,number"`
		Description     string      `field:"description" validate:"required"`
		LongDescription string      `field:"long_description" validate:"required,min=20"`
		Location        null.String `field:"location"`
		IsAvailable     bool        `field:"is_available" validate:"boolean"`
		Realtors        []string    `field:"realtors"`
		Images          []string    `field:"images"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	validate := validator.New()
	if err := validate.Struct(params); err != nil {
		respondWithError(w, 400, fmt.Sprintln(err))
		return
	}

	property, err := apiCgf.DB.CreateProperties(r.Context(), database.CreatePropertiesParams{
		ID:              uuid.New(),
		CreatedAt:       time.Now().UTC(),
		UpdatedAt:       time.Now().UTC(),
		Type:            params.Type,
		Address:         params.Address,
		City:            params.City,
		Bedrooms:        params.Bedrooms,
		Bathrooms:       params.Bathrooms,
		Price:           params.Price,
		SquareMeters:    params.SquareMeters,
		Description:     params.Description,
		LongDescription: params.LongDescription,
		Location:        params.Location,
		IsAvailable:     params.IsAvailable,
		Realtors:        params.Realtors,
		Images:          params.Images,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couln't create user: %v", err))
		return
	}

	respondWithJson(w, 200, databasePropertyToProperty(property))
}
