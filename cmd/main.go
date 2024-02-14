package main

import (
	"net/http"

	"github.com/HarrisonHemstreet/go-ws/internal/handler"
	userHandler "github.com/HarrisonHemstreet/go-ws/internal/handler/user"
	"github.com/HarrisonHemstreet/go-ws/internal/utils"
)

func main() {
	mainMux := http.NewServeMux()

	// Example usage: Set up different routes
	utils.RegisterRoutes(mainMux, "/user", userHandler.UserRouter, []string{"POST"})
	utils.RegisterRoutes(mainMux, "/login", handler.Login, []string{"POST"})

	// Start the HTTP server
	if err := http.ListenAndServe(":8080", mainMux); err != nil {
		panic(err)
	}
}
