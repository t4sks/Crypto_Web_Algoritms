package cipher

import (
	"errors"
	"strings"
)

func PolibiusSquareEncode(line string) (string, error) {
	if len(line) <= 0 {
		return "", errors.New("invalid input")
	}
	line = strings.ToUpper(line)
	chars := []rune(line)
	alphabetEng := "ABCDEFGHIKLMNOPQRSTUVWXYZ"
	alphabetRus := "АБВГДЕЖЗИКЛМНОПРСТУФХЦЧШЩЫЬЭЮЯ"
	var rows []int
	var cols []int

	for _, char := range line {

	}
	return line, nil
}

func PolibiusSquareDecode() {

}
