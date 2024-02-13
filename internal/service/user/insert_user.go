package user

import (
	"github.com/HarrisonHemstreet/go-ws/internal/database"
	"github.com/HarrisonHemstreet/go-ws/internal/model"
)

// InsertUser inserts a new user into the database and returns the inserted user.
func InsertUser(user model.User) (model.User, error) {
	db := database.InitDB()
	defer db.Close()

	statement := `INSERT INTO users (username, email) VALUES ($1, $2) RETURNING id, username, email`
	err := db.QueryRow(statement, user.Username, user.Email).Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
