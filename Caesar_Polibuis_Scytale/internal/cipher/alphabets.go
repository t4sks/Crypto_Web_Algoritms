package cipher

var AlphabetRusUpMatrix = [][]rune{
	{'А', 'Б', 'В', 'Г', 'Д', 'Е'},
	{'Ё', 'Ж', 'З', 'И', 'Й', 'К'},
	{'Л', 'М', 'Н', 'О', 'П', 'Р'},
	{'С', 'Т', 'У', 'Ф', 'Х', 'Ц'},
	{'Ч', 'Ш', 'Щ', 'Ъ', 'Ы', 'Ь'},
	{'Э', 'Ю', 'Я', '.', ',', '!'},
}
var AlphabetEngUPMatrix = [][]rune{
	{'A', 'B', 'C', 'D', 'E'},
	{'F', 'G', 'H', 'I', 'K'},
	{'L', 'M', 'N', 'O', 'P'},
	{'Q', 'R', 'S', 'T', 'U'},
	{'V', 'W', 'X', 'Y', 'Z'},
}

var DictionaryAlphabets = map[string][]rune{
	"AlphabetRussinLowRune": []rune("абвгдеёжзийклмнопрстуфхцчшщъыьэюя"),
	"AlphabetRussinUpRune":  []rune("АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ"),
	"EngAlphabetLowRune":    []rune("abcdefghijklmnopqrstuvwxyz"),
	"EngAlphabetUpRune":     []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ"),
	"SpecialChars":          []rune(`" !\"#$%&'()*+,-./:;<=>?@[\\]^_{|}~"` + "`"),
}
