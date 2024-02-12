package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/HarrisonHemstreet/go-ws/internal/database"
	"github.com/HarrisonHemstreet/go-ws/internal/model"
)

// insertUserHandler handles the HTTP request for inserting a new user
func InsertUser(w http.ResponseWriter, r *http.Request) {
	db := database.InitDB()
	defer db.Close()
	// Only allow POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode the request body into a User struct
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Insert the user into the database
	statement := `INSERT INTO users (username, email) VALUES ($1, $2)`
	_, insert_err := db.Exec(statement, user.Username, user.Email)
	if insert_err != nil {
		http.Error(w, "Failed to insert user", http.StatusInternalServerError)
		return
	}

	// Respond to the client
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User inserted successfully")
}
