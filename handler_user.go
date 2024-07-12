package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Max23strm/real-state-backend/internal/database"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func (apiCgf *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		FirstName    string `field:"first_name" validate:"required,max=32"`
		LastName     string `field:"last_name" validate:"required,max=32"`
		EmailAddress string `field:"email_address"`
		UserName     string `field:"username" validate:"required,max=32"`
		IsAdmin      bool   `field:"is_admin" validate:"boolean"`
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

	user, err := apiCgf.DB.CreateUsers(r.Context(), database.CreateUsersParams{
		ID:           uuid.New(),
		CreatedAt:    time.Now().UTC(),
		UpdatedAt:    time.Now().UTC(),
		FirstName:    params.FirstName,
		LastName:     params.LastName,
		EmailAddress: params.EmailAddress,
		UserName:     params.UserName,
		IsAdmin:      params.IsAdmin,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couln't create user: %v", err))
		return
	}

	respondWithJson(w, 200, databaseUserToUser(user))
}
