package user

import (
	"fmt"

	"github.com/HarrisonHemstreet/go-ws/internal/database"
	"github.com/HarrisonHemstreet/go-ws/internal/model"
	"golang.org/x/crypto/bcrypt"
)

// InsertUser inserts a new user into the database and returns the inserted user.
func InsertUser(user model.User) (model.User, error) {
	db := database.InitDB()
	defer db.Close()

	var hashedPassword []byte
	var err error

	// Ensure Password is not nil before attempting to hash it
	if user.Password != nil {
		hashedPassword, err = bcrypt.GenerateFromPassword([]byte(*user.Password), bcrypt.DefaultCost)
		if err != nil {
			return model.User{}, err
		}
	} else {
		// Handle the case where Password is nil, perhaps default to an error or a specific behavior
		return model.User{}, fmt.Errorf("password is required")
	}

	// Insert the user with the hashed password
	statement := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id, username, email`
	err = db.QueryRow(statement, user.Username, user.Email, string(hashedPassword)).Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		return model.User{}, err
	}

	// Clear the password in the returned user object for security
	user.Password = nil // Correctly set to nil for a *string
	return user, nil
}
