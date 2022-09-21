package task_2

import (
	"errors"
	"fmt"
	"strconv"
)

func unpackingString(str string) (string, error) { // str, err
	var err error
	runes := []rune(str)
	var res string

	for i := 0; i < len(runes); i++ {
		symbol := runes[i]

		_, err2 := strconv.Atoi(string(runes[i]))
		if err2 == nil {
			return "", errors.New("некорректная строка")
		}

		digit := 1
		firstDig := true
		i++
		for err == nil && i < len(runes) { // count of symbols 4, 15, 300 и тд
			newDig, err1 := strconv.Atoi(string(runes[i]))
			if err1 != nil {
				i--
				break
			} else {
				if firstDig {
					firstDig = false
					digit = newDig
				} else {
					digit = digit*10 + newDig
				}
			}
			i++
			err = err1
		}
		err = nil

		for j := 0; j < digit; j++ { // print symbol digit times
			res += string(symbol)
		}

	}

	return res, nil
}

func main() {
	res, ok := unpackingString("")
	fmt.Println("res = ", res, ok)
}
