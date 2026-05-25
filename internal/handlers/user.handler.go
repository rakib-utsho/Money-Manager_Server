package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"money-manager-server/internal/database"
	"money-manager-server/internal/models"
	"money-manager-server/internal/utils"
)

func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriteJSON(
			w,
			http.StatusMethodNotAllowed,
			false,
			"Method not allowed",
			nil,
		)

		return
	}

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		utils.WriteJSON(
			w,
			http.StatusBadRequest,
			false,
			"Invalid request body",
			nil,
		)
		return
	}

	if user.Name == "" || user.Email == "" || user.Password == "" {
		utils.WriteJSON(
			w,
			http.StatusBadRequest,
			false,
			"All fields are required",
			nil,
		)
		return
	}

	query := `
	INSERT INTO users (name, email, password)
	VALUES ($1, $2, $3)
	RETURNING id, created_at, updated_at
	`
	err = database.DB.QueryRow(
		context.Background(),
		query,
		user.Name,
		user.Email,
		user.Password,
	).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		utils.WriteJSON(
			w,
			http.StatusInternalServerError,
			false,
			"Failed to  create user",
			nil,
		)
		return
	}

	utils.WriteJSON(
		w,
		http.StatusCreated,
		true,
		"User registered successfully",
		user,
	)
}
