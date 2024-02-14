package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/HarrisonHemstreet/go-ws/internal/model"
)

func HandleRouteError(w http.ResponseWriter, path string, err error) {
	switch err.Error() {
	case "pq: duplicate key value violates unique constraint \"users_username_key\"":
		routeError(w, path, "A user with that username already exists.", "duplicate_username", http.StatusConflict)
	case "authentication failed: crypto/bcrypt: hashedPassword is not the hash of the given password":
		routeError(w, path, "Authentication Failed: Wrong password or username", "wrong_creds", http.StatusUnauthorized)
	default:
		routeError(w, path, "Unhandled Error", "unhandled_error", http.StatusInternalServerError)
	}
}

func routeError(w http.ResponseWriter, path, message, errorCode string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	errResp := model.ErrorResponse{
		StatusCode:   statusCode,
		ErrorMessage: message,
		ErrorCode:    errorCode,
		Path:         path,
		Timestamp:    time.Now(),
	}
	json.NewEncoder(w).Encode(errResp)
}
