package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	var result strings.Builder
	runes := []rune(str)
	length := len(runes)
	isEscaped := false
	for i := 0; i < length; i++ {
		current := runes[i]
		if current == '\\' {
			if i+1 >= length {
				return "", ErrInvalidString
			}
			next := runes[i+1]
			if next != '\\' && !unicode.IsDigit(next) {
				return "", ErrInvalidString
			}
			result.WriteRune(next)
			isEscaped = true
			i++
			continue
		}
		if unicode.IsDigit(current) {
			if i == 0 || (unicode.IsDigit(runes[i-1]) && !isEscaped) {
				return "", ErrInvalidString
			}
			char := runes[i-1]
			repeat, _ := strconv.Atoi(string(current))
			written := []rune(result.String())
			if len(written) == 0 {
				return "", ErrInvalidString
			}
			if repeat == 0 {
				result.Reset()
				result.WriteString(string(written[:len(written)-1]))
			} else {
				written = written[:len(written)-1]
				result.Reset()
				result.WriteString(string(written))
				result.WriteString(strings.Repeat(string(char), repeat))
			}
			isEscaped = updateEscapeFlag(isEscaped, runes, i)
			continue
		}
		isEscaped = updateEscapeFlag(isEscaped, runes, i)
		result.WriteRune(current)
	}
	// fmt.Printf("=== RESULT: %s\n", result.String())
	return result.String(), nil
}

func updateEscapeFlag(isEscaped bool, runes []rune, i int) bool {
	if isEscaped && i-2 >= 0 && runes[i-2] == '\\' {
		return false
	}
	return isEscaped
}
