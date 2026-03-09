package cipher

import (
	"errors"
	"strings"
	"unicode"
)

func PolybiusSquareEncode(line string, language string) (string, error) {
	if line == "" {
		return "", errors.New("input must not be empty")
	}
	line = strings.TrimSpace(line)
	switch language {
	case "english":
		return englishAlp(line), nil
	case "russian":
		return russianAlp(line), nil
	default:
		return "", errors.New(" unsupported language")
	}
}

func russianAlp(input string) string {
	var row []int
	var columns []int
	chars := []rune(input)
	for i := 0; i < len(chars); i++ {
		upperChar := unicode.ToUpper(chars[i])
		r, c := findIndex(upperChar, AlphabetRusUpMatrix)
		if r != -1 {
			row = append(row, r)
			columns = append(columns, c)
		}
	}
	coordinates := append(row, columns...)

	return rebuild(chars, coordinates, AlphabetRusUpMatrix)
}
func englishAlp(input string) string {
	replace := strings.NewReplacer("j", "i", "J", "I")
	input = replace.Replace(input)
	var row []int
	var columns []int
	chars := []rune(input)
	for i := 0; i < len(chars); i++ {
		upperChar := unicode.ToUpper(chars[i])
		r, c := findIndex(upperChar, AlphabetEngUPMatrix)
		if r != -1 {
			row = append(row, r)
			columns = append(columns, c)
		}
	}
	coordinates := append(row, columns...)

	return rebuild(chars, coordinates, AlphabetEngUPMatrix)
}

func PolybiusSquareDecode(line string, language string) (string, error) {
	if line == "" {
		return "", errors.New("input must not be empty")
	}
	line = strings.TrimSpace(line)
	switch language {
	case "english":
		return englishAlpDecode(line), nil
	case "russian":
		return russianAlpDecode(line), nil
	default:
		return "", errors.New("unsupported language")
	}
}

func russianAlpDecode(input string) string {
	chars := []rune(input)
	var allCoords []int

	for _, char := range chars {
		r, c := findIndex(unicode.ToUpper(char), AlphabetRusUpMatrix)
		if r != -1 {
			allCoords = append(allCoords, r, c)
		}
	}

	if len(allCoords) == 0 {
		return input
	}

	mid := len(allCoords) / 2
	rows := allCoords[:mid]
	cols := allCoords[mid:]

	var decodedCoords []int
	for i := 0; i < len(rows); i++ {
		decodedCoords = append(decodedCoords, rows[i], cols[i])
	}

	return rebuild(chars, decodedCoords, AlphabetRusUpMatrix)
}

func englishAlpDecode(input string) string {
	chars := []rune(input)
	var allCoords []int

	for _, char := range chars {
		r, c := findIndex(unicode.ToUpper(char), AlphabetEngUPMatrix)
		if r != -1 {
			allCoords = append(allCoords, r, c)
		}
	}

	if len(allCoords) == 0 {
		return input
	}

	mid := len(allCoords) / 2
	rows := allCoords[:mid]
	cols := allCoords[mid:]

	var decodedCoords []int
	for i := 0; i < len(rows); i++ {
		decodedCoords = append(decodedCoords, rows[i], cols[i])
	}

	return rebuild(chars, decodedCoords, AlphabetEngUPMatrix)
}

func rebuild(original []rune, coordinatres []int, alphabet [][]rune) string {
	var result strings.Builder
	coordIdx := 0

	for _, r := range original {
		upperR := unicode.ToUpper(r)
		rIdx, _ := findIndex(upperR, alphabet)

		if rIdx == -1 || coordIdx >= len(coordinatres) {
			result.WriteRune(r)
			continue
		}
		newChar := alphabet[coordinatres[coordIdx]][coordinatres[coordIdx+1]]
		coordIdx += 2
		if unicode.IsLower(r) {
			result.WriteRune(unicode.ToLower(newChar))
		} else {
			result.WriteRune(newChar)
		}
	}
	return result.String()
}

func findIndex(char rune, alphabet [][]rune) (int, int) {
	for i, c := range alphabet {
		for j, r := range c {
			if r == char {
				return i, j
			}
		}
	}
	return -1, -1
}
