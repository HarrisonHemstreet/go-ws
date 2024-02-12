package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
)

// CombinedDeleteUserByIDHandler handles the HTTP request for deleting a user by their ID and deletes the user from the database
func DeleteUserByID(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	// Only allow DELETE requests
	if r.Method != http.MethodDelete {
		http.Error(w, "Only DELETE requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	// Assuming the user ID is passed as a URL parameter, e.g., /deleteuserbyid?userid=1
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

	// Execute the DELETE statement to delete the user from the database
	statement := `DELETE FROM users WHERE id = $1`
	result, err := db.Exec(statement, userID)
	if err != nil {
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}
	if rowsAffected == 0 {
		http.Error(w, fmt.Sprintf("No user found with ID %d", userID), http.StatusNotFound)
		return
	}

	// Respond to the client
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User with ID %d deleted successfully", userID)
}
