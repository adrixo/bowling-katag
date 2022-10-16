package src

import "strconv"

func parse(gameResult string) (result int) {
	return
}

func parseChar(input string) (result int, e error) {
	for _, s := range input {
		number, _ := strconv.Atoi(string(s))
		result += number
	}
	return result, e
}
