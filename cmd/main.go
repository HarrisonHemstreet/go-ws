package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/HarrisonHemstreet/go-ws/internal/router"
	"github.com/HarrisonHemstreet/go-ws/internal/utils"
)

func main() {
	mux := router.Router(*utils.Logger)

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
		utils.Logger.Info("Shutting down server...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			utils.Logger.Error("Shutdown error", "error", err)
		} else {
			utils.Logger.Info("Server gracefully stopped")
		}
	}()

	// Start the HTTP server
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		utils.Logger.Error("Server failed to start", "error", err)
		panic(err)
	}
}
