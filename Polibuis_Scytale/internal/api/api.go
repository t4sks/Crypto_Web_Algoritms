package api

import (
	"Polibuis_Scytale/internal/cipher"
	"Polibuis_Scytale/internal/model"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

func apiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	if !strings.HasPrefix(r.Header.Get("Content-Type"), "application/json") {
		http.Error(w, http.StatusText(http.StatusUnsupportedMediaType), http.StatusUnsupportedMediaType)
		return
	}
	r.Body = http.MaxBytesReader(w, r.Body, 10*1024)
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	var request model.Request
	if err := decoder.Decode(&request); err != nil {
		writeError(w, "Invalid Json", http.StatusBadRequest)
		return
	}
	var result string
	var err error
	switch request.Algoritm {
	case "Scytale":
		result, err = scytale(request)
	case "Polibius":
		result, err = polibius(request)
	default:
		writeError(w, "Invalid type of cipher", http.StatusBadRequest)
		return
	}
	if err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(model.Response{Result: result, Error: err})
}

func scytale(requst model.Request) (string, error) {
	var resultScytale string
	var errScytale error
	if requst.Key < 0 && len(requst.Data) > requst.Key {
		return "", errors.New("invalid key, key must shorter than lenght of Data")
	}
	switch requst.Operation {
	case "encrypt":
		resultScytale, errScytale = cipher.Scytale(requst.Data, requst.Key)
		return resultScytale, errScytale
	case "decrypt":
		resultScytale, errScytale = cipher.DecryptScytale(requst.Data, requst.Key)
		return resultScytale, errScytale
	default:
		return "Invalid operation", errors.New(http.StatusText(http.StatusBadRequest))
	}
	return "", errors.New(http.StatusText(http.StatusInternalServerError))
}

func polibius(requst model.Request) (string, error) {
	var resultPolibius string
	var errPolibius error
	switch requst.Operation {
	case "encrypt":
		resultPolibius, errPolibius = cipher.PolibiusSquareEncode(requst.Data, requst.Language)
		return resultPolibius, errPolibius
	case "decrypt":
		resultPolibius, errPolibius = cipher.PolibiusSquareDecode(requst.Data, requst.Language)
		return resultPolibius, errPolibius
	default:
		return "Invalid operation", errors.New(http.StatusText(http.StatusBadRequest))

	}
	return "", errors.New(http.StatusText(http.StatusInternalServerError))
}

func writeError(w http.ResponseWriter, msg string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(model.ErrorResponse{Error: msg})

}
