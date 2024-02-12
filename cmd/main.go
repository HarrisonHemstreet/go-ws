package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/HarrisonHemstreet/go-ws/internal/database"
	"github.com/HarrisonHemstreet/go-ws/internal/handler"
	"github.com/HarrisonHemstreet/go-ws/internal/handler/hello"
	"github.com/HarrisonHemstreet/go-ws/internal/model"
	_ "github.com/lib/pq"
)

func main() {
	db := database.InitDB()
	defer db.Close()

	// Setting up the HTTP server
	http.HandleFunc("/", hello.HelloWorldHandler)
	http.HandleFunc("/insertuser", func(w http.ResponseWriter, r *http.Request) {
		handler.InsertUser(db, w, r)
	})
	http.HandleFunc("/fetchuserbyid", func(w http.ResponseWriter, r *http.Request) {
		handler.FetchUserByID(db, w, r)
	})
	http.HandleFunc("/deleteuserbyid", func(w http.ResponseWriter, r *http.Request) {
		DeleteUserByIDHandler(db, w, r)
	})
	http.HandleFunc("/updateuserbyid", func(w http.ResponseWriter, r *http.Request) {
		UpdateUserByIDHandler(db, w, r)
	})

	fmt.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

// DeleteUserByID deletes a user from the database by user_id
func DeleteUserByID(db *sql.DB, userID int) error {
	// Prepare the DELETE statement
	statement := `DELETE FROM users WHERE id = $1`
	result, err := db.Exec(statement, userID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no user found with ID %d", userID)
	}

	return nil
}

// DeleteUserByIDHandler handles the HTTP request for deleting a user by their ID
func DeleteUserByIDHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	// Only allow DELETE requests
	if r.Method != http.MethodDelete {
		http.Error(w, "Only DELETE requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	// Assuming the user ID is passed as a URL parameter, e.g., /deleteuserbyid?userid=1
	keys, ok := r.URL.Query()["id"]
	if !ok || len(keys[0]) < 1 {
		http.Error(w, "User ID must be provided as a query parameter", http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(keys[0])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	err = DeleteUserByID(db, userID)
	if err != nil {
		if err.Error() == fmt.Sprintf("no user found with ID %d", userID) {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		}
		return
	}

	// Respond to the client
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User with ID %d deleted successfully", userID)
}

// UpdateUserByID updates user details in the database by user_id
func UpdateUserByID(db *sql.DB, userID int, username, email string) error {
	// Prepare the UPDATE statement
	statement := `UPDATE users SET username = $2, email = $3 WHERE id = $1`
	result, err := db.Exec(statement, userID, username, email)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no user found with ID %d", userID)
	}

	return nil
}

// UpdateUserByIDHandler handles the HTTP request for updating a user by their ID
func UpdateUserByIDHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	// Only allow PUT requests
	if r.Method != http.MethodPut {
		http.Error(w, "Only PUT requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	// Assuming the user ID is passed as a URL parameter, e.g., /updateuserbyid?userid=1
	keys, ok := r.URL.Query()["id"]
	if !ok || len(keys[0]) < 1 {
		http.Error(w, "User ID must be provided as a query parameter", http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(keys[0])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Decode the request body for new user details
	var user model.User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = UpdateUserByID(db, userID, user.Username, user.Email)
	if err != nil {
		if err.Error() == fmt.Sprintf("no user found with ID %d", userID) {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, "Failed to update user", http.StatusInternalServerError)
		}
		return
	}

	// Respond to the client
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User with ID %d updated successfully", userID)
}
