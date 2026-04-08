package httpserver

import (
	"Crypto-Ciphers/internal/cipher"
	"Crypto-Ciphers/internal/model"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

func handleApi(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, r, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	if !strings.HasPrefix(r.Header.Get("Content-Type"), "application/json") {
		writeError(w, r, http.StatusText(http.StatusUnsupportedMediaType), http.StatusUnsupportedMediaType)
		return
	}
	r.Body = http.MaxBytesReader(w, r.Body, 10*1024)
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	var request model.Request
	if err := decoder.Decode(&request); err != nil {
		writeError(w, r, "Invalid json", http.StatusBadRequest)
		return
	}
	result, err, code := executeCipher(request)
	if err != nil {
		writeError(w, r, err.Error(), http.StatusBadRequest)
		return
	}
	writeSuccess(w, r, result, code)
}

func executeCipher(request model.Request) (string, error, string) {
	switch strings.ToLower(strings.TrimSpace(request.Algorithm)) {
	case "scytale":
		result, err := cipher.ExecuteScytale(request)
		return result, err, ""
	case "polibius":
		result, err := cipher.ExecutePolibius(request)
		return result, err, ""
	case "caesar":
		result, err := cipher.ExecuteCaesar(request)
		return result, err, ""
	case "cardano":
		result, code, err := cipher.ExecuteCaradan(request)
		return result, err, code
	case "gronsfeld":
		result, err := cipher.ExecuteGronsfeld(request)
		return result, err, ""
	case "vigenere":
		result, err := cipher.ExecuteVigener(request)
		return result, err, ""
	default:
		return "", errors.New("Invalid algorithm"), ""
	}
}
