package main

import (
	"fmt"
	"net/http"

	"github.com/HarrisonHemstreet/go-ws/internal/database"
	"github.com/HarrisonHemstreet/go-ws/internal/handler"
	"github.com/HarrisonHemstreet/go-ws/internal/handler/hello"
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
		handler.DeleteUserByID(db, w, r)
	})
	http.HandleFunc("/updateuserbyid", func(w http.ResponseWriter, r *http.Request) {
		handler.UpdateUserByID(db, w, r)
	})

	fmt.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
