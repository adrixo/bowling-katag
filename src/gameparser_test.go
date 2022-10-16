package src

import (
	"testing"
)

func Test_parse_number(t *testing.T) {
	numbers := []string{"18", "9", "27", "36", "45"}
	expecteds := []int{9, 9, 9, 9, 9}
	for i, number := range numbers {
		expected := expecteds[i]
		result, _ := parseChar(number)
		if expected != result {
			t.Errorf("Expected %v; but got %v", expected, result)
		}
	}
}
