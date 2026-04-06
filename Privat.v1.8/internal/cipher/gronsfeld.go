package cipher

import "strings"

func GronsfeldEncrypt(key string, line string) (string, string) {
	return gronsfeldProcess(key, line, 1)
}

func GronsfeldDecrypt(key string, line string) (string, string) {
	return gronsfeldProcess(key, line, -1)
}

func gronsfeldProcess(key string, line string, direction int) (string, string) {
	textRunes := []rune(line)
	if len(textRunes) == 0 {
		return "", "Строка для шифра не должна быть пуста"
	}

	keyRunes := []rune(key)
	if len(keyRunes) == 0 {
		return "", "Ключ не должен быть пустым"
	}

	for _, r := range keyRunes {
		if r < '0' || r > '9' {
			return "", "Ключ должен состоят только из цифр"
		}
	}
	var result strings.Builder
	result.Grow(len(textRunes))

	keyIndex := 0
	for _, r := range textRunes {
		shift := int(keyRunes[keyIndex] - '0')
		shift *= direction

		newRune, found := shiftRune(r, shift)
		if found {
			result.WriteRune(newRune)
		} else {
			result.WriteRune(r)
		}

		keyIndex++
		if keyIndex >= len(keyRunes) {
			keyIndex = 0
		}
	}
	return result.String(), ""
}

func shiftRune(r rune, shift int) (rune, bool) {
	for _, alphabet := range DictionaryAlphabets {
		for i, char := range alphabet {
			if r == char {
				newIndex := (i + shift) % len(alphabet)
				if newIndex < 0 {
					newIndex += len(alphabet)
				}
				newRune := alphabet[newIndex]
				return newRune, true
			}
		}
	}
	return r, false
}
