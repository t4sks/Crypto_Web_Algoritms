package httpserver

import (
	"Polibuis_Scytale/internal/httpserver/middleware"
	"Polibuis_Scytale/internal/model"
	"encoding/json"
	"net/http"
)

func writeSuccess(w http.ResponseWriter, r *http.Request, result string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_ = json.NewEncoder(w).Encode(model.SuccessResponse{
		Result:    result,
		RequestId: middleware.GetRequestID(r.Context()),
	})
}

func writeError(w http.ResponseWriter, r *http.Request, msg string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(model.ErrorResponse{
		Error:     msg,
		RequestId: middleware.GetRequestID(r.Context()),
	})
}
