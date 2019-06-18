package week2

import (
	"errors"
	"strings"
)

var numbers = map[rune]int{
	'0': 0,
	'1': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
}

func UnpackString(s string) (string, error) {
	var nextChar rune = -1
	var prevChar rune = -1
	multiplier := 0
	escapeNext := false

	// пишем итоговую строку в Builder
	var sb strings.Builder

	for _, ch := range []rune(s) {
		if escapeNext {
			nextChar = ch
			escapeNext = false
		} else if ch == '\\' {
			escapeNext = true
		} else if n, ok := numbers[ch]; ok {
			if multiplier < 1 && n < 1 {
				return "", errors.New("multiplier must be greater than 0")
			}
			if prevChar < 0 {
				return "", errors.New("no character to unpack")
			}
			multiplier = multiplier*10 + n
			continue

		} else {
			nextChar = ch
		}

		if prevChar >= 0 && nextChar >= 0 {
			writeToBuilder(prevChar, multiplier, &sb)
			multiplier = 0
		}
		if nextChar >= 0 {
			prevChar = nextChar
			nextChar = -1
		}
	}
	if prevChar >= 0 {
		writeToBuilder(prevChar, multiplier, &sb)
	}
	return sb.String(), nil
}

// вспомогательная функция для многократной записи символа в буфер
func writeToBuilder(c rune, mul int, builder *strings.Builder) {
	if mul == 0 {
		builder.WriteRune(c)
		return
	}

	for j := 0; j < mul; j++ {
		builder.WriteRune(c)
	}
}
