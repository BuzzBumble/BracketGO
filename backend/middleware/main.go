package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

type StatusTracker struct {
	http.ResponseWriter
	Status int
}

func (r *StatusTracker) WriteHeader(status int) {
	r.Status = status
	r.ResponseWriter.WriteHeader(status)
}

func RequestLoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		tracker := &StatusTracker{
			ResponseWriter: w,
			Status:         http.StatusOK, // Default if WriteHeader not called
		}

		next.ServeHTTP(tracker, r)

		slog.Info("Request",
			slog.String("Method", r.Method),
			slog.String("URL", r.URL.String()),
			slog.Int("Status", tracker.Status),
			slog.Int64("Duration", time.Since(startTime).Milliseconds()),
		)

	})
}

func EnableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*") // Allow any origin
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Check if the request is for CORS preflight
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Pass down the request to the next middleware (or final handler)
		next.ServeHTTP(w, r)
	})
}

func JSONContentTypeMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set JSON Content-Type
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
