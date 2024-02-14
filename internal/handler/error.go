package handler

import (
	"encoding/json"
	"net/http"

	"github.com/HarrisonHemstreet/go-ws/internal/model"
)

func HandleRouteError(w http.ResponseWriter, err error) {
	switch err.Error() {
	case "pq: duplicate key value violates unique constraint \"users_username_key\"":
		Error(w, "A user with that username already exists.", "duplicate_username", http.StatusConflict)
	case "authentication failed: crypto/bcrypt: hashedPassword is not the hash of the given password":
		Error(w, "Authentication Failed: Wrong password or username", "wrong_creds", http.StatusUnauthorized)
	default:
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Error(w http.ResponseWriter, message, errorCode string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	errResp := model.ErrorResponse{
		StatusCode:   statusCode,
		ErrorMessage: message,
		ErrorCode:    errorCode,
	}
	json.NewEncoder(w).Encode(errResp)
}
