package api

import (
	"encoding/json"
	"net/http"
	"strings"
	"Polybuis_sqare+Scytale/internal/model.go"
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
	var

}
