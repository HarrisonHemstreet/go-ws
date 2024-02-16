package user

import (
	"database/sql"
	"fmt"

	"github.com/HarrisonHemstreet/go-ws/internal/database"
	"github.com/HarrisonHemstreet/go-ws/internal/model"
)

// UpdateUser updates a user's details in the database by their ID.
func UpdateUser(userID int, user model.User) (model.User, error) {
	db := database.InitDB()
	defer db.Close()

	statement := `UPDATE users SET username = $2, email = $3 WHERE id = $1 RETURNING id, username, email`
	err := db.QueryRow(statement, userID, user.Username, user.Email).Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.User{}, fmt.Errorf("update unsuccessful")
		}
		return model.User{}, err
	}

	return user, nil
}
