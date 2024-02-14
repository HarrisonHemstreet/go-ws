package handler

import (
	"encoding/json"
	"net/http"

	"github.com/HarrisonHemstreet/go-ws/internal/model"
)

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
