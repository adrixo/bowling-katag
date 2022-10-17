package src

import (
	"strconv"
	"strings"
)

func parseGame(gameResult string) (result int) {
	gameParts := strings.Split(gameResult, "||")
	result = parse(gameParts[0])
	extraResult := parse(gameParts[1])

	return result + extraResult

}

func parse(gameResult string) (result int) {
	frames := strings.Split(gameResult, "|")
	doubler := 0
	for _, frame := range frames {
		frameResult := 0
		for _, symbol := range frame {
			switch string(symbol) {
			case "-":
				doubler--
			case "X":
				frameResult = 10
				if doubler >= 1 {
					frameResult = 20
				}
				doubler = 2
				break
			case "/":
				frameResult = 10
				if doubler >= 1 {
					frameResult = 20
				}
				doubler = 1
				break
			default:
				value, _ := strconv.Atoi(string(symbol))
				frameResult += value
				if doubler >= 1 {
					doubler--
					frameResult += value
				}
			}
		}
		result += frameResult
	}
	return result
}
