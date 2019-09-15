package encoder

import (
	"errors"
	"fmt"
	"strconv"
)

func isNumeric(number rune) bool {
	return number >= '0' && number <= '9'
}

func processEscapeSequence(escapeSequence []rune) ([]rune, bool) {
	var newRune = []rune{}
	last := 1
	backSlashCount := 0
	for escapeSequence[len(escapeSequence)-last] == '\\' {
		backSlashCount++
		last++
	}

	newRune = escapeSequence[0 : len(escapeSequence)-last+1]
	for i := 0; i < backSlashCount/2; i++ {
		newRune = append(newRune, '\\')
	}

	fmt.Printf("~~%s \n", string(newRune))
	return newRune, backSlashCount%2 != 0
}

func EncodeString(str string) (string, error) {
	var newRune = []rune{}

	fmt.Printf("%s \n", str)
	for i, c := range str {
		if !isNumeric(c) {
			newRune = append(newRune, c)
			continue
		}

		if i == 0 {
			return "", errors.New("first character can't be number")
		}

		var needEscape bool
		newRune, needEscape = processEscapeSequence(newRune)

		if i > 0 && needEscape {
			newRune = append(newRune, c)
			continue
		}

		count, _ := strconv.Atoi(string(c))
		for i := 1; i < count; i++ {
			newRune = append(newRune, newRune[len(newRune)-1])
		}
	}
	fmt.Printf("%s \n", string(newRune))

	return string(newRune), nil
}
