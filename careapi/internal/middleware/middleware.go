package middleware

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type Middleware func(http.Handler) http.Handler

type contextKey string

const RequestIDKey contextKey = "requestID"

func Chain(h http.Handler, middlewares ...Middleware) http.Handler{
	for i := len(middlewares) - 1; i >= 0; i--{
		h = middlewares[i](h)
	}
	return h
}

func RequestID(next http.Handler) http.Handler{
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request){
		id := r.Header.Get("X-Request-ID")

		if id == ""{
			id = uuid.New().String()
		}

		ctx := context.WithValue(r.Context(), RequestIDKey, id)
		w.Header().Set("X-Request-ID", id)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func Logger(log *slog.Logger) Middleware{
	return func(next http.Handler) http.Handler{
		return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request){
			start := time.Now()

			rw := &responseWriter{ResponseWriter: w, status: http.StatusOK}

			next.ServeHTTP(rw, r)

			requestID, _ := r.Context().Value(RequestIDKey).(string)
			log.Info("request", 
		"requestId", requestID,
	"method", r.Method,
"path", r.URL.Path,
"status", rw.status,
"latency", time.Since(start).String())
		})
	}
}


// responseWriter wraps the standard http.ResponseWriter to capture the status code.
// Because the standard library doesn't let you "read" the status after it's written,
// this spy intercepts the write, saves the number, and then passes it along.
type responseWriter struct {
    http.ResponseWriter // This means it inherits all standard methods from w
    status int          // Our secret spy variable
}

// WriteHeader is called by the endpoint (e.g., when it does w.WriteHeader(http.StatusNotFound))
func (rw *responseWriter) WriteHeader(status int) {
    rw.status = status                  // 1. Save the status code for our Logger
    rw.ResponseWriter.WriteHeader(status) // 2. Actually send the status to the user
}