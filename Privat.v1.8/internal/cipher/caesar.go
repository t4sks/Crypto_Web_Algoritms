package cipher

func CaesarEncrypt(key int, line string) (string, string) {
	chars := []rune(line)
	var encoded []rune
	for _, char := range chars {
		alphabet, index := cesarFindIndex(char)
		alpha := DictionaryAlphabets[alphabet]
		if index != -1 {
			index = (index + (key % len(alpha))) % len(alpha)
			encoded = append(encoded, DictionaryAlphabets[alphabet][index])
		} else {
			encoded = append(encoded, char)
		}
	}
	return string(encoded), ""
}

func CaesarDecrypt(key int, line string) (string, string) {
	chars := []rune(line)
	var decoded []rune
	for _, char := range chars {
		alphabet, index := cesarFindIndex(char)
		alpha := DictionaryAlphabets[alphabet]
		if index != -1 {
			index = (index - (key % len(alpha)) + len(alpha)) % len(alpha)
			decoded = append(decoded, DictionaryAlphabets[alphabet][index])
		} else {
			decoded = append(decoded, char)
		}
	}
	return string(decoded), ""
}

func cesarFindIndex(char rune) (string, int) {
	for key, alphabet := range DictionaryAlphabets {
		for i, r := range alphabet {
			if r == char {
				return key, i
			}
		}
	}
	return "", -1
}
