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
	//escapeNext := false

	var sb strings.Builder
	for i := 0; i < len(s); i++ {
		ch := rune(s[i])
		if nextChar >= 0 {
			prevChar = nextChar
			nextChar = -1
		}
		if n, ok := numbers[ch]; ok {
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
	}
	if nextChar >= 0 {
		writeToBuilder(nextChar, multiplier, &sb)

	}
	return sb.String(), nil
}

func writeToBuilder(c rune, mul int, builder *strings.Builder) {
	if mul == 0 {
		mul = 1
	}
	for j := 0; j < mul; j++ {
		builder.WriteRune(c)
	}
}
