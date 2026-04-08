package cipher

func CaesarEncrypt(key int, line string) (string, string) {
	chars := []rune(line)
	var encoded []rune
	for _, char := range chars {
		alphabet, index := FindIndex(char)
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
		alphabet, index := FindIndex(char)
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
