package service

import (
	"net/http"

	"github.com/HarrisonHemstreet/go-ws/internal/handler"
)

func HandleRouteError(w http.ResponseWriter, err error) {
	switch err.Error() {
	case "pq: duplicate key value violates unique constraint \"users_username_key\"":
		handler.Error(w, "A user with that username already exists.", "duplicate_username", http.StatusConflict)
	default:
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
