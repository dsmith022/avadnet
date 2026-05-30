package api

import (
	"avadnet/backend/internal/api/models"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (app *application) registerUserHandler(w http.ResponseWriter, r *http.Request) {
	var req models.CreateUserRequest

	if err := app.readJSON(w, r, &req); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	now := time.Now()
	resp := models.UserResponse{
		ID:        uuid.New(),
		Username:  req.Username,
		Email:     req.Email,
		CreatedAt: now,
		UpdatedAt: now,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) loginUserHandler(w http.ResponseWriter, r *http.Request) {
	app.notImplementedResponse(w, r, "login not implemented")
}
