package handler

import (
	"encoding/json"
	"net/http"

	service "github.com/HarrisonHemstreet/go-ws/internal/service/user"
	"github.com/HarrisonHemstreet/go-ws/internal/utils/jwt"
)

func Login(w http.ResponseWriter, r *http.Request) {
	// Decode the request body
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Authenticate the user
	if service.AuthenticateUser(credentials.Username, credentials.Password) {
		// Generate a token
		tokenString, err := jwt.CreateToken(credentials.Username)
		if err != nil {
			http.Error(w, "Failed to generate token", http.StatusInternalServerError)
			return
		}

		// Respond with the token
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
	} else {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
	}
}
