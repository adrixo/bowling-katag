package src

import "testing"

func Test_parse_number(t *testing.T) {
	numbers := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	expecteds := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i, number := range numbers {
		expected := expecteds[i]
		result, _ := parseChar(number)
		if expected != result {
			t.Errorf("Expected %v; but got %v", expected, result)
		}
	}
}
