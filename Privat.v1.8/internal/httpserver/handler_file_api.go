package httpserver

import (
	"Crypto-Ciphers/internal/model"
	"bytes"
	"errors"
	"io"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"unicode/utf8"
)

const (
	maxTextFileSize    = 10 * 1024
	maxUploadSize      = 64 * 1024
	maxMultipartmemory = 128 * 1024
)

var (
	allowedFileExtensions = map[string]struct{}{
		".txt": {},
	}

	suspiciousMiddleExtensions = map[string]struct{}{
		".html": {},
		".php":  {},
		".js":   {},
		".exe":  {},
		".sh":   {},
		".bat":  {},
		".cmd":  {},
	}
)

func handleFileProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, r, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	if !strings.HasPrefix(r.Header.Get("Content-Type"), "multipart/form-data") {
		writeError(w, r, http.StatusText(http.StatusUnsupportedMediaType), http.StatusUnsupportedMediaType)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, maxMultipartmemory+maxUploadSize)

	if err := r.ParseMultipartForm(maxMultipartmemory); err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "file too big") {
			writeError(w, r, "Файл слишком большой", http.StatusRequestEntityTooLarge)
			return
		}
		writeError(w, r, "некорректная форма загрузки файла", http.StatusUnsupportedMediaType)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		writeError(w, r, "Файл не был загружен", http.StatusBadRequest)
		return
	}
	defer file.Close()

	if err := validateUploadedFileName(header.Filename); err != nil {
		writeError(w, r, err.Error(), http.StatusBadRequest)
		return
	}

	data, err := io.ReadAll(file)
	if err != nil {
		writeError(w, r, "Ошибка чтения файла", http.StatusBadRequest)
		return
	}

	if len(data) == 0 {
		writeError(w, r, "Пустой файл", http.StatusBadRequest)
		return
	}

	if len(data) > maxTextFileSize {
		writeError(w, r, "Файл слишком большой", http.StatusRequestEntityTooLarge)
		return
	}

	data = bytes.TrimPrefix(data, []byte{0xEF, 0xBB, 0xBF})

	if !utf8.Valid(data) {
		writeError(w, r, "Файл должен содержать текст в формате UTF-8", http.StatusBadRequest)
		return
	}

	detectMIME := http.DetectContentType(data)
	if !strings.HasPrefix(detectMIME, "text/plain") {
		writeError(w, r, "Содержимое файла не текст", http.StatusBadRequest)
		return
	}

	if bytes.Contains(data, []byte{0}) {
		writeError(w, r, "Обнаружены недопустимые бинарные символы", http.StatusBadRequest)
		return
	}

	request, err := buildFileRequest(r, string(data))
	if err != nil {
		writeError(w, r, err.Error(), http.StatusBadRequest)
		return
	}

	result, execError, code := executeCipher(request)

	if execError != nil {
		writeError(w, r, execError.Error(), http.StatusBadRequest)
		return
	}

	safeFileName := buildSafeResultFileName(header.Filename)

	job := fileStore.Save(processedFileJob{
		OriginName:   header.Filename,
		SafeFileName: safeFileName,
		Algorithm:    request.Algorithm,
		Operation:    request.Operation,
		ResultText:   result,
		Code:         code,
	})
	writeFileSuccess(w, r, result, code, job.ID)
}

func validateUploadedFileName(filename string) error {
	cleanName := filepath.Base(strings.TrimSpace(filename))
	if cleanName == "" || cleanName == "." || cleanName == "/" {
		return errors.New("Некорректное имя файла")
	}
	cleanName = filepath.Clean(cleanName)

	ext := strings.ToLower(filepath.Ext(cleanName))
	if _, ok := allowedFileExtensions[ext]; !ok {
		return errors.New("Неподдерживаемый тип файла")
	}

	nameWithoutExt := strings.TrimSuffix(cleanName, ext)
	innerExt := strings.ToLower(filepath.Ext(nameWithoutExt))
	if _, suspicious := suspiciousMiddleExtensions[innerExt]; suspicious {
		return errors.New("Двойное расширение")
	}
	return nil
}

func buildFileRequest(r *http.Request, data string) (model.Request, error) {
	keyRaw := strings.TrimSpace(r.FormValue("key"))
	key := 0
	if keyRaw != "" {
		parsedKey, err := strconv.Atoi(keyRaw)
		if err != nil {
			return model.Request{}, errors.New("Некорретный ключ")
		}
		key = parsedKey
	}

	keyStr := strings.TrimSpace(r.FormValue("keyString"))

	request := model.Request{
		Algorithm: strings.TrimSpace(r.FormValue("algorithm")),
		Data:      data,
		Language:  strings.TrimSpace(r.FormValue("language")),
		Operation: strings.TrimSpace(r.FormValue("operation")),
		Key:       key,
		Code:      strings.TrimSpace(r.FormValue("code")),
		KeyString: keyStr,
	}

	if request.Algorithm == "" {
		return model.Request{}, errors.New("Алгоритм не выбран")
	}
	if request.Operation == "" {
		return model.Request{}, errors.New("Операция не выбрана")
	}

	return request, nil
}

func buildSafeResultFileName(originalName string) string {
	base := filepath.Base(strings.TrimSpace(originalName))
	ext := strings.ToLower(filepath.Ext(base))
	nameWithoutExt := strings.TrimSuffix(base, ext)
	newExtension := ".txt"

	var newName strings.Builder

	newName.WriteString(nameWithoutExt + newExtension)
	return newName.String()
}
