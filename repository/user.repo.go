package repository

import (
	"database/sql"

	"money-manager-server/models"
)

// UserRepository holds the db connection
// all user related SQL queries live here
type UserRepository struct {
	DB *sql.DB
}

// CreateUser inserts a new user into the database
// password here is already hashed — hashing happens in service layer
func (r *UserRepository) CreateUser(user *models.User) error {
	query := `
	INSERT INTO users (name, email, password)
	VALUES ($1, $2, $3)
	RETURNING id, created_at
	`
	// $1, $2, $3 are placeholders — prevents SQL injection
	// RETURNING gives us back the auto-generated id and created_at
	return r.DB.QueryRow(query, user.Name, user.Email, user.Password).Scan(&user.ID, &user.CreatedAt)
}

func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	query := `
	SELECT id, name, email, password, created_at
	FORM users
	WHERE email = $1
	`
	user := &models.User{}

	err := r.DB.QueryRow(query, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil // no user found with that email
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}
