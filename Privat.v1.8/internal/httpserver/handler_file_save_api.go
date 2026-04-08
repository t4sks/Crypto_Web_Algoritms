package httpserver

import (
	"bytes"
	"net/http"
	"strconv"
	"strings"
)

func handleFileSave(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, r, "Method not allowed", http.StatusBadRequest)
		return
	}
	prefix := "/api/file/download/"
	if !strings.HasPrefix(r.URL.Path, prefix) {
		writeError(w, r, "Not found", http.StatusNotFound)
		return
	}

	fileId := strings.TrimPrefix(r.URL.Path, "/api/file/download/")

	if fileId == "" {
		writeError(w, r, "Bad file", http.StatusBadRequest)
		return
	}

	file, ok := fileStore.Get(fileId)
	if !ok {
		writeError(w, r, "File not found", http.StatusNotFound)
		return
	}

	filename := file.SafeFileName

	w.Header().Set("Content-Type", `text/plain; charset=utf-8`)
	w.Header().Set("Content-Disposition", `attachment; filename="`+filename+`"`)
	w.Header().Set("Content-Length", strconv.Itoa(len(file.ResultText)))

	http.ServeContent(w, r, filename, file.CreateTime, bytes.NewReader([]byte(file.ResultText)))
}
