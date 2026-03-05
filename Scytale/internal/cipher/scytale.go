package cipher

import (
	"errors"
)

var (
	ErrInvalidKey     = errors.New("Ключ m должен быть больше 0 и меньше длины сообщения")
	ErrInvalidMessage = errors.New("Сообщение слишком короткое")
)

func Scytale(line string, m int) (string, error) {
	arrayOfline := []rune(line)
	k := len(arrayOfline)
	if m <= 0 || m >= k {
		return "", ErrInvalidKey
	}
	if k == 0 {
		return "", ErrInvalidMessage
	}
	n := ((k - 1) / m) + 1
	//fmt.Printf("k=%d, m=%d, n=%d\n", k, m, n)
	newline := make([]rune, m*n)
	for i := range newline {
		newline[i] = ' '
	}
	for i, letter := range arrayOfline {
		index := (m * (i % n)) + (i / n)
		if index < len(newline) {
			newline[index] = letter
		}
	}
	return string(newline), nil
}

func DecryptScytale(line string, m int) (string, error) {
	if m <= 0 {
		return "", ErrInvalidKey
	}
	arrayOfline := []rune(line)
	k := len(arrayOfline)
	if k == 0 {
		return "", ErrInvalidMessage
	}
	n := ((k - 1) / m) + 1
	decoded := make([]rune, k)
	for i, letter := range arrayOfline {
		index := (n * (i % m)) + (i / m)
		if index < k {
			decoded[index] = letter
		}
	}
	return string(decoded), nil
}
