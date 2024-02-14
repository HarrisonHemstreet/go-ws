package user

import (
	"encoding/json"
	"net/http"

	"github.com/HarrisonHemstreet/go-ws/internal/handler"
	"github.com/HarrisonHemstreet/go-ws/internal/model"
	userService "github.com/HarrisonHemstreet/go-ws/internal/service/user"
)

// InsertUser handles the HTTP request for inserting a new user.
func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	insertedUser, err := userService.InsertUser(user)
	if err != nil {
		handler.HandleRouteError(w, r.URL.Path, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(insertedUser)
}
