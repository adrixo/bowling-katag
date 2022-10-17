package src

import (
	"strconv"
	"strings"
)

func parse(gameResult string) (result int) {
	frames := strings.Split(gameResult, "|")
	for _, frame := range frames {
		result += parseSymbol(string(frame[0]))
		result += parseSymbol(string(frame[1]))
	}
	return result
}

func parseSymbol(symbol string) (result int) {
	if symbol == "-" {
		return 0
	}
	result, _ = strconv.Atoi(symbol)
	return result
}
