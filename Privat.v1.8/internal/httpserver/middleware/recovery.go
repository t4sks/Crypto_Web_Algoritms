package middleware

import (
	"encoding/json"
	"log"
	"net/http"
)

func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				log.Printf("request_id=%s, panic=%v", GetRequestID(r.Context()), rec)
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.WriteHeader(http.StatusInternalServerError)

				_ = json.NewEncoder(w).Encode(map[string]string{
					"error":      "Internal Server Error",
					"request_id": GetRequestID(r.Context()),
				})
			}
		}()
		next.ServeHTTP(w, r)
	})
}
