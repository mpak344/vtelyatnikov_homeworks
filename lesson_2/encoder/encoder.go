package encoder

import (
	"fmt"
	"strconv"
	"errors"
)

func isNumeric(number rune) bool {
	return number >= '0' && number <= '9'
}

func EncodeString(str string) (string, error) {
	var newRune = []rune{}

	fmt.Printf("%s \n", str)
	for i, c := range str {
		if !isNumeric(c) {
			newRune = append(newRune, c)
			continue
		}
		if i > 0  && newRune[len(newRune)-1] == '\\'{
			newRune = append(newRune[0: len(newRune) - 1], c)
			continue
		}
		if i == 0 {
			return "", errors.New("first character can't be number")
		}
		count, _ := strconv.Atoi(string(c))
		for i := 1; i < count; i++ {
			newRune = append(newRune, newRune[len(newRune)-1])
		}
	}
	fmt.Printf("%s \n", string(newRune))

	return string(newRune), nil
}
