package user

import (
	"fmt"

	"github.com/HarrisonHemstreet/go-ws/internal/database"
	"github.com/HarrisonHemstreet/go-ws/internal/model"
)

// UpdateUser updates a user's details in the database by their ID.
func UpdateUser(userID int, user model.User) error {
	db := database.InitDB()
	defer db.Close()

	statement := `UPDATE users SET username = $2, email = $3 WHERE id = $1`
	result, err := db.Exec(statement, userID, user.Username, user.Email)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error checking affected rows: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no user found with ID %d", userID)
	}

	return nil
}
