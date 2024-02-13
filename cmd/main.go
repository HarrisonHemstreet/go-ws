package main

import (
	"net/http"

	handler "github.com/HarrisonHemstreet/go-ws/internal/handler/user"
)

func main() {
	http.HandleFunc("/user", handler.UserRouter)

	// Start the HTTP server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
