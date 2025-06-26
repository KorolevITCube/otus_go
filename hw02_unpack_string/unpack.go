package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	var sb strings.Builder
	runes := []rune(str)
	for i := 0; i < len(runes); i++ {
		switch {
		case string(runes[i]) == `\` && i != (len(runes)-1):
			i++
			fallthrough
		case unicode.IsLetter(runes[i]):
			sb.WriteString(multiple(runes, i))
		case unicode.IsDigit(runes[i]) && i != (len(runes)-1) && (unicode.IsDigit(runes[i+1]) || i == 0):
			return "", ErrInvalidString
		}
	}
	finalString := sb.String()
	return finalString, nil
}

func multiple(data []rune, i int) string {
	if i != (len(data)-1) && unicode.IsDigit(data[i+1]) {
		return strings.Repeat(string(data[i]), int(data[i+1]-'0'))
	}
	return string(data[i])
}
