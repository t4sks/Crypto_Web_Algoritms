package cipher

import (
	"Crypto-Ciphers/internal/model"
	"errors"
	"strings"
	"unicode/utf8"
)

func ExecuteScytale(request model.Request) (string, error) {
	if request.Key <= 0 {
		return "", errors.New("invalid key, key must be more than 0")
	}
	if utf8.RuneCountInString(request.Data) < request.Key {
		return "", errors.New("invalid key, key must shorter than lenght of message")
	}
	switch strings.ToLower(strings.TrimSpace(request.Operation)) {
	case "encrypt":
		return Scytale(request.Data, request.Key)
	case "decrypt":
		return DecryptScytale(request.Data, request.Key)
	default:
		return "", errors.New("invalid operation")
	}
}

func ExecutePolibius(request model.Request) (string, error) {
	switch strings.ToLower(strings.TrimSpace(request.Operation)) {
	case "encrypt":
		return PolybiusSquareEncode(request.Data, request.Language)
	case "decrypt":
		return PolybiusSquareDecode(request.Data, request.Language)
	default:
		return "", errors.New("invalid operation")
	}
}

func ExecuteCaesar(request model.Request) (string, error) {
	if request.Key <= 0 {
		return "", errors.New("invalid key, key must be more than 0")
	}
	switch strings.ToLower(strings.TrimSpace(request.Operation)) {
	case "encrypt":
		result, err := CaesarEncrypt(request.Key, request.Data)
		if err != "" {
			return "", errors.New(err)
		}
		return result, nil
	case "decrypt":
		result, err := CaesarDecrypt(request.Key, request.Data)
		if err != "" {
			return "", errors.New(err)
		}
		return result, nil
	default:
		return "", errors.New("Invalid operation")
	}
}

func ExecuteCaradan(request model.Request) (string, string, error) {
	if request.Key <= 0 {
		return "", "", errors.New("invalid key, key must be more than 0")
	}
	switch strings.ToLower(strings.TrimSpace(request.Operation)) {
	case "encrypt":
		result, code, err := CaradanEncrypt(request.Key, request.Data)
		if err != "" {
			return "", "", errors.New(err)
		}
		return result, code, nil
	case "decrypt":
		result, err := CardanoDecrypt(request.Key, request.Data, request.Code)
		if err != "" {
			return "", "", errors.New(err)
		}
		return result, "", nil
	default:
		return "", "", errors.New("Invalid operation")
	}
}

func ExecuteGronsfeld(request model.Request) (string, error) {
	if len(request.KeyString) <= 0 {
		return "", errors.New("Ключ не должен быть пустым")
	}
	switch strings.ToLower(strings.TrimSpace(request.Operation)) {
	case "encrypt":
		result, err := GronsfeldEncrypt(request.KeyString, request.Data)
		if err != "" {
			return "", errors.New(err)
		}
		return result, nil
	case "decrypt":
		result, err := GronsfeldDecrypt(request.KeyString, request.Data)
		if err != "" {
			return "", errors.New(err)
		}
		return result, nil
	default:
		return "", errors.New("Invalid operation")
	}
}

func ExecuteVigener(request model.Request) (string, error) {
	if len(request.KeyString) <= 0 {
		return "", errors.New("invalid key, key must be longer than 0")
	}
	switch strings.ToLower(strings.TrimSpace(request.Operation)) {
	case "encrypt":
		result, err := VigenereEncrypt(request.KeyString, request.Data)
		if err != "" {
			return "", errors.New(err)
		}
		return result, nil
	case "decrypt":
		result, err := VigenereDecrypt(request.KeyString, request.Data)
		if err != "" {
			return "", errors.New(err)
		}
		return result, nil
	default:
		return "", errors.New("Invalid operation")
	}
}
