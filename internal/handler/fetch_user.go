package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/HarrisonHemstreet/go-ws/internal/database"
	"github.com/HarrisonHemstreet/go-ws/internal/model"
)

// CombinedFetchUserByIDHandler handles the HTTP request for fetching a user by their ID and fetches the user from the database
func FetchUserByID(w http.ResponseWriter, r *http.Request) {
	db := database.InitDB()
	defer db.Close()
	// Only allow GET requests
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	// Assuming the user ID is passed as a URL parameter, e.g., /fetchuserbyid?id=1
	keys, ok := r.URL.Query()["id"]
	if !ok || len(keys[0]) < 1 {
		http.Error(w, "User ID must be provided as a query parameter", http.StatusBadRequest)
		return
	}

	ID, err := strconv.Atoi(keys[0])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Fetch user from the database by user_id
	var user model.User
	err = db.QueryRow("SELECT id, username, email FROM users WHERE id = $1", ID).Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		fmt.Println(err)
		if err == sql.ErrNoRows {
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to fetch user", http.StatusInternalServerError)
		}
		return
	}

	// Marshal the User struct to JSON
	jsonResponse, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Error converting user data to JSON", http.StatusInternalServerError)
		return
	}

	// Write the JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
