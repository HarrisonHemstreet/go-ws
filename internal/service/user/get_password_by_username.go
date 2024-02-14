package user

import (
	"github.com/HarrisonHemstreet/go-ws/internal/database"
)

// GetPasswordByUsername retrieves a user's password hash from the database by their username.
func GetPasswordByUsername(username string) (string, error) {
	db := database.InitDB()
	defer db.Close()

	var password string
	err := db.QueryRow("SELECT password FROM users WHERE username = $1", username).Scan(&password)
	if err != nil {
		return "", err
	}

	return password, nil
}
