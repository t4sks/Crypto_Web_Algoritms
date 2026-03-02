package cipher

import (
	"errors"
	"strings"
	"unicode"
)

var alphabetRusUp = [][]rune{
	{'А', 'Б', 'В', 'Г', 'Д', 'Е'},
	{'Ж', 'З', 'И', 'К', 'Л', 'М'},
	{'Н', 'О', 'П', 'Р', 'С', 'Т'},
	{'У', 'Ф', 'Х', 'Ц', 'Ч', 'Ш'},
	{'Щ', 'Ы', 'Ь', 'Э', 'Ю', 'Я'},
}
var alphabetEngUP = [][]rune{
	{'A', 'B', 'C', 'D', 'E'},
	{'F', 'G', 'H', 'I', 'K'},
	{'L', 'M', 'N', 'O', 'P'},
	{'Q', 'R', 'S', 'T', 'U'},
	{'V', 'W', 'X', 'Y', 'Z'},
}

func PolibiusSquareEncode(line string, language string) (string, error) {
	if line == "" {
		return "", errors.New("input must not be empty")
	}
	line = strings.TrimSpace(line)
	switch language {
	case "English":
		return englishAlp(line), nil
	case "Russian":
		return russianAlp(line), nil
	default:
		return "", errors.New("unsupported language")
	}
	return "", nil
}

func russianAlp(input string) string {
	replace := strings.NewReplacer("Ё", "Е", "ё", "е", "Й", "И", "й", "и", "Ъ", "Ь", "ъ", "ь")
	input = replace.Replace(input)
	var row []int
	var columns []int
	chars := []rune(input)
	for i := 0; i < len(chars); i++ {
		upperChar := unicode.ToUpper(chars[i])
		r, c := findIndex(upperChar, alphabetRusUp)
		if r != -1 {
			row = append(row, r)
			columns = append(columns, c)
		}
	}
	coordinates := append(row, columns...)

	return rebuild(chars, coordinates, alphabetRusUp)
}
func englishAlp(input string) string {
	replace := strings.NewReplacer("j", "i", "J", "I")
	input = replace.Replace(input)
	var row []int
	var columns []int
	chars := []rune(input)
	for i := 0; i < len(chars); i++ {
		upperChar := unicode.ToUpper(chars[i])
		r, c := findIndex(upperChar, alphabetEngUP)
		if r != -1 {
			row = append(row, r)
			columns = append(columns, c)
		}
	}
	coordinates := append(row, columns...)

	return rebuild(chars, coordinates, alphabetEngUP)
}

func PolibiusSquareDecode(line string, language string) (string, error) {
	if line == "" {
		return "", errors.New("input must not be empty")
	}
	line = strings.TrimSpace(line)
	switch language {
	case "English":
		return englishAlpDecode(line), nil
	case "Russian":
		return russianAlpDecode(line), nil
	default:
		return "", errors.New("unsupported language")
	}
}

func russianAlpDecode(input string) string {
	chars := []rune(input)
	var allCoords []int

	for _, char := range chars {
		r, c := findIndex(unicode.ToUpper(char), alphabetRusUp)
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

	return rebuild(chars, decodedCoords, alphabetRusUp)
}

func englishAlpDecode(input string) string {
	chars := []rune(input)
	var allCoords []int

	for _, char := range chars {
		r, c := findIndex(unicode.ToUpper(char), alphabetEngUP)
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

	return rebuild(chars, decodedCoords, alphabetEngUP)
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
