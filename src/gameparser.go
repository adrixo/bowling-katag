package src

import (
	"strconv"
	"strings"
)

func parse(gameResult string) (result int) {
	frames := strings.Split(gameResult, "|")
	for _, frame := range frames {
		frameResult := 0
		for _, symbol := range frame {
			if string(symbol) == "-" {
				continue
			}
			if string(symbol) == "X" {
				frameResult = 10
				continue
			}
			if string(symbol) == "/" {
				frameResult = 10
				continue
			}
			value, _ := strconv.Atoi(string(symbol))
			frameResult += value
		}
		result += frameResult
	}
	return result
}
