package main

import (
	"net/http"

	"github.com/HarrisonHemstreet/go-ws/internal/handler"
	"github.com/HarrisonHemstreet/go-ws/internal/handler/user"
	"github.com/HarrisonHemstreet/go-ws/internal/middleware/jwt"
)

func main() {
	mux := http.NewServeMux()

	// Middleware
	jwtMiddleware := jwt.ValidateToken

	// Login
	mux.HandleFunc("POST /api/v1/login", handler.Login)

	// User routes
	userRoutes := http.NewServeMux()
	userRoutes.HandleFunc("GET /api/v1/user", user.GetUser)
	userRoutes.HandleFunc("POST /api/v1/user", user.CreateUser)
	userRoutes.HandleFunc("PUT /api/v1/user", user.UpdateUser)
	userRoutes.HandleFunc("DELETE /api/v1/user", user.DeleteUser)
	mux.Handle("/api/v1/user", jwtMiddleware(userRoutes))

	// Start the HTTP server
	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}
