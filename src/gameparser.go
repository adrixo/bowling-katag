package src

import (
	"strconv"
	"strings"
)

func parse(gameResult string) (result int) {
	frames := strings.Split(gameResult, "|")
	for _, frame := range frames {
		for _, symbol := range frame {
			number, _ := strconv.Atoi(string(symbol))
			result += number
		}
	}
	return result
}
