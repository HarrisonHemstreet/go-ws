package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/HarrisonHemstreet/go-ws/internal/model"
	"github.com/HarrisonHemstreet/go-ws/internal/utils"
)

func HandleRouteError(w http.ResponseWriter, path string, err error) {
	switch err.Error() {
	case "pq: duplicate key value violates unique constraint \"users_username_key\"":
		RouteError(w, path, "A user with that username already exists.", "duplicate_username", http.StatusConflict)
	case "authentication failed: crypto/bcrypt: hashedPassword is not the hash of the given password":
		RouteError(w, path, "Authentication Failed: Wrong password or username", "wrong_creds", http.StatusUnauthorized)
	case "update unsuccessful":
		RouteError(w, path, "Update unsuccessful.", "update_unsuccessful", http.StatusUnprocessableEntity)
	default:
		RouteError(w, path, "Unhandled Error", "unhandled_error", http.StatusInternalServerError)
	}
}

func RouteError(w http.ResponseWriter, path, message, errorCode string, statusCode int) {
	utils.Logger.Error(message, errorCode, statusCode)
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
