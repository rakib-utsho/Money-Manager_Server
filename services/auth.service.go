package services

import (
	"errors"
	"money-manager-server/models"
	"money-manager-server/repository"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepo *repository.UserRepository
}

// Register hashes the password then saves the user
// business logic lives here, not in handler or repository
func (s *AuthService) Register(name, email, password string) (*models.User, error) {

	// check if email already exists
	existing, err := s.UserRepo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return nil, errors.New("email already registered")
	}

	// bcrypt hashes the password with a random salt
	// cost 12 means it runs 2^12 hashing rounds — slow enough to resist brute force
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword), // store hash not plain text
	}

	// save to database via repository
	if err := s.UserRepo.CreateUser(user); err != nil {
		return nil, err
	}

	return user, nil
}

// Login verifies credentials and returns a JWT token
func (s *AuthService) Login(email, password string) (string, error) {

	// fetch user from database
	user, err := s.UserRepo.GetUserByEmail(email)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", errors.New("invalid email or password")
	}

	// CompareHashAndPassword checks if the plain password matches the hash
	// returns error if they don't match
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	// generate JWT token
	token, err := generateJWT(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

// generateJWT creates a signed token containing the user's ID
// this token is sent to client and used for all protected requests
func generateJWT(userID int) (string, error) {

	// claims are the data stored inside the token
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(72 * time.Hour).Unix(), // token expires in 72 hours
	}

	// create token with HS256 signing method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// sign the token with our secret key from .env
	// anyone with this secret can verify the token is genuine
	secret := os.Getenv("JWT_SECRET")
	return token.SignedString([]byte(secret))
}
