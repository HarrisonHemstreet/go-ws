package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/HarrisonHemstreet/go-ws/internal/handler"
	"github.com/HarrisonHemstreet/go-ws/internal/handler/user"
	"github.com/HarrisonHemstreet/go-ws/internal/middleware"
	"github.com/HarrisonHemstreet/go-ws/internal/middleware/jwt"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))

	mux := http.NewServeMux()

	jwtMiddleware := jwt.ValidateToken

	loggedMux := middleware.HttpLogger(mux, logger)

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
	if err := http.ListenAndServe(":8080", loggedMux); err != nil {
		panic(err)
	}
}
