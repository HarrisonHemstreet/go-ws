package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

func HttpLogger(handler http.Handler, logger *slog.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		rec := &responseRecorder{ResponseWriter: w, statusCode: http.StatusOK}
		handler.ServeHTTP(rec, r)
		duration := time.Since(startTime)

		// Log using slog in JSON format
		logger.Info("received a HTTP request",
			"method", r.Method,
			"path", r.RequestURI,
			"status_code", rec.statusCode,
			"status_text", http.StatusText(rec.statusCode),
			"duration", duration,
		)
	})
}

type responseRecorder struct {
	http.ResponseWriter
	statusCode int
}

func (rec *responseRecorder) WriteHeader(statusCode int) {
	rec.statusCode = statusCode
	rec.ResponseWriter.WriteHeader(statusCode)
}
