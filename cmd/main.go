package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/HarrisonHemstreet/go-ws/internal/router"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))
	mux := router.Router(*logger)

	// Initialize server with ability to shutdown gracefully
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
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
