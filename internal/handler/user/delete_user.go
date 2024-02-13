package handler

import (
	"fmt"
	"net/http"
	"strconv"

	service "github.com/HarrisonHemstreet/go-ws/internal/service/user"
)

// DeleteUserByID handles the HTTP request for deleting a user by their ID
func DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Only DELETE requests are allowed", http.StatusMethodNotAllowed)
		return
	}

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

	err = service.DeleteUserByID(userID)
	if err != nil {
		if err == service.ErrUserNotFound {
			http.Error(w, fmt.Sprintf("No user found with ID %d", userID), http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User with ID %d deleted successfully", userID)
}
