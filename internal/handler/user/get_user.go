package user

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/HarrisonHemstreet/go-ws/internal/handler"
	service "github.com/HarrisonHemstreet/go-ws/internal/service/user"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	keys, ok := r.URL.Query()["id"]
	if !ok || len(keys[0]) < 1 {
		http.Error(w, "User ID must be provided as a query parameter", http.StatusBadRequest)
		return
	}

	ID, err := strconv.Atoi(keys[0])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := service.FetchUserByID(ID)
	if err != nil {
		handler.HandleRouteError(w, r.URL.Path, err)
		return
	}

	jsonResponse, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Error converting user data to JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
