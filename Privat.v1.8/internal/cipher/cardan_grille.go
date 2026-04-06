package cipher

import (
	"math/rand"
	"strconv"
	"strings"
)

type Point struct {
	Row int
	Col int
}

func CaradanEncrypt(k int, line string) (string, string, string) {
	if k <= 0 {
		return "", "", "Значение K должно быть больше 0"
	}
	runes := []rune(line)
	if len(runes) <= 0 {
		return "", "", "Строка шифрования пуста"
	}
	code := makeCode(k)
	blocksize := 4 * k * k
	holes, _ := buildCardanoHoles(k, code)

	for len(runes)%blocksize != 0 {
		runes = append(runes, '~')
	}

	var result strings.Builder
	for i := 0; i < len(runes); i += blocksize {
		result.WriteString(encryptCardanoBlock(runes[i:i+blocksize], k, holes))
	}

	var codeResult strings.Builder
	for i := 0; i < len(code); i++ {
		codeResult.WriteString(strconv.Itoa(code[i]))
	}

	return result.String(), codeResult.String(), ""
}

func buildCardanoHoles(k int, code []int) ([]Point, string) {

	n := 2 * k
	holes := make([]Point, 0, k*k)

	for i := 0; i < k*k; i++ {
		r := i / k
		c := i % k

		switch code[i] {
		case 1:
			holes = append(holes, Point{Row: r, Col: c})
		case 2:
			holes = append(holes, Point{Row: c, Col: n - 1 - r})
		case 3:
			holes = append(holes, Point{Row: n - 1 - r, Col: n - 1 - c})
		case 4:
			holes = append(holes, Point{Row: n - 1 - c, Col: r})
		default:
			return nil, "неверный код"
		}
	}
	return holes, ""
}

func makeCode(k int) []int {
	code := make([]int, k*k)
	for i := 0; i < k*k; i++ {
		code[i] = rand.Intn(4) + 1
	}
	return code
}

func encryptCardanoBlock(block []rune, k int, holes []Point) string {
	n := 2 * k
	grid := make([][]rune, n)
	for i := range grid {
		grid[i] = make([]rune, n)
	}

	current := make([]Point, len(holes))
	copy(current, holes)

	partSize := k * k
	offset := 0

	for turn := 0; turn < 4; turn++ {
		for i, p := range current {
			grid[p.Row][p.Col] = block[offset+i]
		}
		offset += partSize

		for i := range current {
			current[i] = rotatePointCW(current[i], n)
		}
	}

	var b strings.Builder
	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			b.WriteRune(grid[r][c])
		}
	}

	return b.String()
}

func CardanoDecrypt(k int, text string, code string) (string, string) {
	runes := []rune(text)
	if len(runes) == 0 {
		return "", "Ввод не должен быть пустым"
	}

	blockSize := 4 * k * k
	if len(runes)%blockSize != 0 {
		return "", "Неверная длина"
	}

	codeRunes := []rune(code)
	codes := make([]int, len(codeRunes))

	for i := 0; i < len(codeRunes); i++ {
		if codeRunes[i] < '1' || codeRunes[i] > '4' {
			return "", "неверный код"
		}
		codes[i] = int(codeRunes[i] - '0')
	}

	holes, err := buildCardanoHoles(k, codes)
	if err != "" {
		return "", err
	}

	var b strings.Builder
	for i := 0; i < len(runes); i += blockSize {
		b.WriteString(decryptCardanoBlock(runes[i:i+blockSize], k, holes))
	}

	return b.String(), ""
}

func rotatePointCW(p Point, size int) Point {
	return Point{
		Row: p.Col,
		Col: size - 1 - p.Row,
	}
}

func decryptCardanoBlock(block []rune, k int, holes []Point) string {
	n := 2 * k
	grid := make([][]rune, n)
	for i := range grid {
		grid[i] = make([]rune, n)
	}

	idx := 0
	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			grid[r][c] = block[idx]
			idx++
		}
	}

	current := make([]Point, len(holes))
	copy(current, holes)

	var b strings.Builder
	for turn := 0; turn < 4; turn++ {
		for _, p := range current {
			b.WriteRune(grid[p.Row][p.Col])
		}
		for i := range current {
			current[i] = rotatePointCW(current[i], n)
		}
	}

	return b.String()
}
