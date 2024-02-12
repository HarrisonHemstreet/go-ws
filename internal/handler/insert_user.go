package handler

import (
	"encoding/json"
	"net/http"

	"github.com/HarrisonHemstreet/go-ws/internal/database"
	"github.com/HarrisonHemstreet/go-ws/internal/model"
)

// InsertUser handles the HTTP request for inserting a new user
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

	// Insert the user into the database and return the inserted user
	statement := `INSERT INTO users (username, email) VALUES ($1, $2) RETURNING id, username, email`
	// Assuming the database setup includes an auto-generated ID for the users table
	err = db.QueryRow(statement, user.Username, user.Email).Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		http.Error(w, "Failed to insert user", http.StatusInternalServerError)
		return
	}

	// Respond to the client with the inserted user
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
