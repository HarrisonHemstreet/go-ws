package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword takes a plain text password and returns a bcrypt hashed password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
