package cipher

func VigenereEncrypt(key string, line string) (string, string) {
	lineChars := []rune(line)
	keyChars := []rune(key)

	if len(keyChars) == 0 {
		return "", "ключ пуст"
	}

	var encoded []rune
	keyIndex := 0

	for _, char := range lineChars {
		alphabetLine, indexLine := FindIndex(char)

		if indexLine == -1 {
			encoded = append(encoded, char)
			continue
		}

		alphaLine := DictionaryAlphabets[alphabetLine]

		_, indexKey := FindIndex(keyChars[keyIndex])
		if indexKey == -1 {
			indexKey = 0
		}

		newIndexLine := (indexLine + indexKey) % len(alphaLine)
		encoded = append(encoded, alphaLine[newIndexLine])

		keyIndex = (keyIndex + 1) % len(keyChars)
	}

	return string(encoded), ""
}

func VigenereDecrypt(key string, line string) (string, string) {
	lineChars := []rune(line)
	keyChars := []rune(key)

	if len(keyChars) == 0 {
		return "", "ключ пуст"
	}

	var decoded []rune
	keyIndex := 0

	for _, char := range lineChars {
		alphabetLine, indexLine := FindIndex(char)

		if indexLine == -1 {
			decoded = append(decoded, char)
			continue
		}

		alphaLine := DictionaryAlphabets[alphabetLine]

		_, indexKey := FindIndex(keyChars[keyIndex])
		if indexKey == -1 {
			indexKey = 0
		}

		newIndexLine := (indexLine - indexKey + len(alphaLine)) % len(alphaLine)
		decoded = append(decoded, alphaLine[newIndexLine])

		keyIndex = (keyIndex + 1) % len(keyChars)
	}

	return string(decoded), ""
}
