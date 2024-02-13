package handler

import (
	"net/http"
)

func UserRouter(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		FetchUserByID(w, r)
	case http.MethodPost:
		InsertUser(w, r)
	case http.MethodPut:
		UpdateUserByID(w, r)
	case http.MethodDelete:
		DeleteUserByID(w, r)
	default:
		http.Error(w, "Unsupported HTTP method", http.StatusMethodNotAllowed)
	}
}
