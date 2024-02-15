package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	// Log server starting message
	logger.Info("Starting server on port 8080")

	// Initialize server with ability to shutdown gracefully
	server := &http.Server{
		Addr:    ":8080",
		Handler: loggedMux,
	}

	// Listen for interrupt signal to gracefully shutdown the server
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-stopChan // wait for interrupt signal
		logger.Info("Shutting down server...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			logger.Error("Shutdown error", "error", err)
		} else {
			logger.Info("Server gracefully stopped")
		}
	}()

	// Start the HTTP server
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Error("Server failed to start", "error", err)
		panic(err)
	}
}
