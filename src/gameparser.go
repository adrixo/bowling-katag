package src

import (
	"strconv"
	"strings"
)

func parse(gameResult string) (result int) {
	frames := strings.Split(gameResult, "|")
	for _, frame := range frames {
		firstThrow, _ := strconv.Atoi(string(frame[0]))
		secondThrow, _ := strconv.Atoi(string(frame[1]))
		result += firstThrow
		result += secondThrow
	}
	return result
}
