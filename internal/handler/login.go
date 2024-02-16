package handler

import (
	"encoding/json"
	"net/http"

	"github.com/HarrisonHemstreet/go-ws/internal/service/user"
)

func Login(w http.ResponseWriter, r *http.Request) {
	// Decode the request body
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		RouteError(w, r.URL.Path, "Invalid request body", "bad_token", http.StatusUnauthorized)
		return
	}

	// Authenticate the user
	tokenString, err := user.AuthenticateUser(credentials.Username, credentials.Password)
	if err != nil {
		HandleRouteError(w, r.URL.Path, err)
		return
	}

	// Respond with the token if authentication is successful
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}
