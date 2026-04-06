package middleware

import (
	"log"
	"net/http"
	"time"
)

type statusWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *statusWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		sw := &statusWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}
		next.ServeHTTP(sw, r)
		log.Printf("request_id=%s method=%s url=%s status_code=%d duration=%s remote_addr=%s",
			GetRequestID(r.Context()),
			r.Method,
			r.URL.Path,
			sw.statusCode,
			time.Since(start).Round(time.Millisecond),
			r.RemoteAddr)
	})
}
