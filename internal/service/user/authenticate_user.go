package user

import (
	"fmt"

	"github.com/HarrisonHemstreet/go-ws/internal/utils/jwt"
	"golang.org/x/crypto/bcrypt"
)

// AuthenticateUser checks the provided username and password against the database.
func AuthenticateUser(username, password string) (string, error) {
	// Retrieve user from database based on username
	// For demonstration, let's assume you have a function getUserByUsername that returns a user and an error
	passwordHash, err := GetPasswordByUsername(username)
	if err != nil {
		// Handle error (e.g., user not found)
		return "", fmt.Errorf("authentication failed: %w", err)
	}

	// Now compare the provided password with the one stored in the database
	err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	if err != nil {
		// Passwords do not match
		return "", fmt.Errorf("authentication failed: %w", err)
	}

	// If passwords match, generate JWT token
	tokenString, err := jwt.CreateToken(username)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}

	return tokenString, nil
}
