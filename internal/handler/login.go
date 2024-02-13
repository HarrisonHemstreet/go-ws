package handler

import (
	"encoding/json"
	"net/http"

	"github.com/HarrisonHemstreet/go-ws/internal/utils/jwt"
)

// Assume a mock function for user authentication
func authenticateUser(username, password string) bool {
	// Implement user authentication against database or in-memory store
	// For simplicity, this is a placeholder function
	return true // Mock authentication: always true
}

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
	if authenticateUser(credentials.Username, credentials.Password) {
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
