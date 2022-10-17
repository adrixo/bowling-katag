package src

import (
	"strconv"
	"strings"
)

func parse(gameResult string) (result int) {
	frames := strings.Split(gameResult, "|")
	for _, frame := range frames {
		firstSymbol := string(frame[0])
		var firstThrow int
		if firstSymbol == "-" {
			firstThrow = 0
		} else {
			firstThrow, _ = strconv.Atoi(firstSymbol)
		}
		result += firstThrow

		secondSymbol := string(frame[1])
		var secondThrow int
		if secondSymbol == "-" {
			secondThrow = 0
		} else {
			secondThrow, _ = strconv.Atoi(secondSymbol)
		}
		result += secondThrow
	}
	return result
}
