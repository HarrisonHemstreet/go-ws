package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/HarrisonHemstreet/go-ws/internal/model"
)

// CombinedUpdateUserByIDHandler handles the HTTP request for updating a user by their ID and updates the user details in the database
func UpdateUserByID(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	// Only allow PUT requests
	if r.Method != http.MethodPut {
		http.Error(w, "Only PUT requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	// Assuming the user ID is passed as a URL parameter, e.g., /updateuserbyid?userid=1
	keys, ok := r.URL.Query()["id"]
	if !ok || len(keys[0]) < 1 {
		http.Error(w, "User ID must be provided as a query parameter", http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(keys[0])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Decode the request body for new user details
	var user model.User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Prepare and execute the UPDATE statement to update the user in the database
	statement := `UPDATE users SET username = $2, email = $3 WHERE id = $1`
	result, err := db.Exec(statement, userID, user.Username, user.Email)
	if err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}
	if rowsAffected == 0 {
		http.Error(w, fmt.Sprintf("No user found with ID %d", userID), http.StatusNotFound)
		return
	}

	// Respond to the client
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User with ID %d updated successfully", userID)
}
