package httpserver

import (
	"Polibuis_Scytale/internal/cipher"
	"Polibuis_Scytale/internal/model"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"unicode/utf8"
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
		writeError(w, r, "invalid json", http.StatusBadRequest)
		return
	}
	result, err := executeCipher(request)
	if err != nil {
		writeError(w, r, err.Error(), http.StatusBadRequest)
		return
	}
	writeSuccess(w, r, result)
}

func executeCipher(request model.Request) (string, error) {
	switch request.Algoritm {
	case "Scytale":
		return executeScytale(request)
	case "Polibius":
		return executePolibius(request)
	case "Caesar":
		return executeCaesar(request)
	default:
		return "", errors.New("invalid algoritm")
	}
}

func executeScytale(request model.Request) (string, error) {
	if request.Key <= 0 {
		return "", errors.New("invalid key, key must be more than 0")
	}
	if utf8.RuneCountInString(request.Data) < request.Key {
		return "", errors.New("invalid key, key must shorter than lenght of message")
	}
	switch request.Operation {
	case "encrypt":
		return cipher.Scytale(request.Data, request.Key)
	case "decrypt":
		return cipher.DecryptScytale(request.Data, request.Key)
	default:
		return "", errors.New("invalid operation")
	}
}

func executePolibius(request model.Request) (string, error) {
	switch request.Operation {
	case "encrypt":
		return cipher.PolybiusSquareEncode(request.Data, request.Language)
	case "decrypt":
		return cipher.PolybiusSquareDecode(request.Data, request.Language)
	default:
		return "", errors.New("invalid operation")
	}
}

func executeCaesar(request model.Request) (string, error) {
	if request.Key <= 0 {
		return "", errors.New("invalid key, key must be more than 0")
	}
	switch request.Operation {
	case "encrypt":
		result, _ := cipher.CaesarEncrypt(request.Key, request.Data)
		return result, nil
	case "decrypt":
		result, _ := cipher.CaesarDecrypt(request.Key, request.Data)
		return result, nil
	default:
		return "", errors.New("invalid operation")
	}
}
