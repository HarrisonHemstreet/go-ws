package user

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/HarrisonHemstreet/go-ws/internal/handler"
	service "github.com/HarrisonHemstreet/go-ws/internal/service/user"
)

// DeleteUserByID handles the HTTP request for deleting a user by their ID
func DeleteUser(w http.ResponseWriter, r *http.Request) {
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
		handler.HandleRouteError(w, r.URL.Path, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User with ID %d deleted successfully", userID)
}
