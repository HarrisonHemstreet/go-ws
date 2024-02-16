package user

import (
	"net/http"
	"strconv"

	"github.com/HarrisonHemstreet/go-ws/internal/handler"
	service "github.com/HarrisonHemstreet/go-ws/internal/service/user"
)

// DeleteUserByID handles the HTTP request for deleting a user by their ID
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["id"]
	if !ok || len(keys[0]) < 1 {
		handler.RouteError(w, r.URL.Path, "User ID must be provided as a query parameter", "missing_user_id", http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(keys[0])
	if err != nil {
		handler.RouteError(w, r.URL.Path, "Invalid user ID", "missing_user_id", http.StatusBadRequest)
		return
	}

	err = service.DeleteUserByID(userID)
	if err != nil {
		handler.HandleRouteError(w, r.URL.Path, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
