package service

import (
	"errors"

	"github.com/HarrisonHemstreet/go-ws/internal/database"
)

var ErrUserNotFound = errors.New("user not found")

// DeleteUserByID deletes a user from the database by their ID
func DeleteUserByID(userID int) error {
	db := database.InitDB()
	defer db.Close()

	statement := `DELETE FROM users WHERE id = $1`
	result, err := db.Exec(statement, userID)
	if err != nil {
		return err // Could wrap in a custom error type if needed
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err // Could wrap in a custom error type if needed
	}
	if rowsAffected == 0 {
		return ErrUserNotFound
	}

	return nil
}
