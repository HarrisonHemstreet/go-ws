package service

import (
	"github.com/HarrisonHemstreet/go-ws/internal/database"
	"github.com/HarrisonHemstreet/go-ws/internal/model"
)

// FetchUserByID retrieves a user from the database by their ID.
func FetchUserByID(ID int) (model.User, error) {
	db := database.InitDB()
	defer db.Close()

	var user model.User
	err := db.QueryRow("SELECT id, username, email FROM users WHERE id = $1", ID).Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
