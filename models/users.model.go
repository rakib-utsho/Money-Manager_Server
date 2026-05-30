package models

import "time"

// User mirrors the "users" table in database
// json tags control what field name appears in API response
type User struct {
	ID        int       `json: "id"`
	Name      string    `json: "name"`
	Email     string    `json: "email"`
	Password  string    `json: "password"`
	CreatedAt time.Time `json: "created_at"`
}
