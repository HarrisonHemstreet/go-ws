package user

import (
	"net/http"
)

func UserRouter(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetUser(w, r)
	case http.MethodPost:
		CreateUser(w, r)
	case http.MethodPut:
		UpdateUserByID(w, r)
	case http.MethodDelete:
		DeleteUserByID(w, r)
	default:
		http.Error(w, "Unsupported HTTP method", http.StatusMethodNotAllowed)
	}
}
