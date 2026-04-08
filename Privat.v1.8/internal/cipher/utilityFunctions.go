package cipher

func FindIndex(char rune) (string, int) {
	for key, alphabet := range DictionaryAlphabets {
		for i, r := range alphabet {
			if r == char {
				return key, i
			}
		}
	}
	return "", -1
}
