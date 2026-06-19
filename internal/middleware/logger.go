// Package middleware provides common middleware for the application, such as logging, authentication, etc.
package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

type responseWrite struct {
	http.ResponseWriter
	status int
	wrote  bool
}

func (rw *responseWrite) WriteHeader(status int) {
	if rw.wrote {
		return
	}

	rw.status = status
	rw.wrote = true
	rw.ResponseWriter.WriteHeader(status)
}

func Logger(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			timeStart := time.Now()
			rw := &responseWrite{ResponseWriter: w}

			next.ServeHTTP(rw, r)

			duration := time.Since(timeStart)

			level := slog.LevelInfo

			if rw.status >= 500 {
				level = slog.LevelError
			} else if rw.status >= 400 {
				level = slog.LevelWarn
			}

			logger.LogAttrs(r.Context(), level, "HTTP Request",
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
				slog.Int("status", rw.status),
				slog.String("duration", duration.String()),
				// slog.String("remote_addr", r.RemoteAddr),
				// slog.String("user_agent", r.UserAgent()),
			)
		})

	}
}
