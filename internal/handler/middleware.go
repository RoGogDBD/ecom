package handler

import (
	"log"
	"net/http"
	"time"
)

type Middleware func(http.Handler) http.Handler

// responseWriter оборачивает http.ResponseWriter для захвата статус-кода и размера ответа.
type responseWriter struct {
	http.ResponseWriter
	statusCode int
	size       int
}

// newResponseWriter создает новый responseWriter с дефолтным статусом 200.
func newResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{
		ResponseWriter: w,
		statusCode:     http.StatusOK,
	}
}

// WriteHeader перехватывает статус-код.
func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

// Write перехватывает размер ответа.
func (rw *responseWriter) Write(b []byte) (int, error) {
	size, err := rw.ResponseWriter.Write(b)
	rw.size += size
	return size, err
}

func Conveyor(h http.Handler, middlewares ...Middleware) http.Handler {
	for _, middleware := range middlewares {
		h = middleware(h)
	}

	return h
}

func LoggingMiddleware(logger *log.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			start := time.Now()
			wrapped := newResponseWriter(w)

			next.ServeHTTP(wrapped, req)

			duration := time.Since(start)
			logger.Printf(
				"%s %s %d %s %dB %s",
				req.Method,
				req.URL.Path,
				wrapped.statusCode,
				req.RemoteAddr,
				wrapped.size,
				duration,
			)
		})
	}
}
