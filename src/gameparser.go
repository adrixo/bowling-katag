package src

import "strconv"

func parse(gameResult string) (result int) {
	return
}

func parseFrame(input string) (result int, e error) {
	for _, s := range input {
		if string(s) == "-" {
			continue
		}
		number, _ := strconv.Atoi(string(s))
		result += number
	}
	return result, e
}
