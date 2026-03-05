package handler

import (
	"encoding/json"
	"net/http"
	"scytale/internal/cipher"
	"scytale/internal/model"
)

func ScytaleHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Content-Type not supported", http.StatusUnsupportedMediaType)
		return
	}
	r.Body = http.MaxBytesReader(w, r.Body, 10*1024)
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	var req model.CipherRequest
	if err := decoder.Decode(&req); err != nil {
		writeError(w, "Invalid json", http.StatusBadRequest)
		return
	}
	if req.Key <= 0 || req.Key >= len(req.Text) {
		writeError(w, "Invalid key, Or length of message", http.StatusBadRequest)
		return
	}

	var result string
	var err error
	switch req.Operation {
	case "encrypt":
		result, err = cipher.Scytale(req.Text, req.Key)
	case "decrypt":
		result, err = cipher.DecryptScytale(req.Text, req.Key)
	default:
		writeError(w, "Invalid operation", http.StatusBadRequest)
		return
	}
	if err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(model.EncryptResponse{Result: result})
}

func writeError(w http.ResponseWriter, msg string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(model.ErrorResponse{
		Error: msg,
	})
}
