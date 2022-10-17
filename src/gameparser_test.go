package src

import (
	"testing"
)

func Test_parse_number_frame(t *testing.T) {
	numbers := []string{"18", "9", "27", "36", "45"}
	expecteds := []int{9, 9, 9, 9, 9}
	for i, number := range numbers {
		expected := expecteds[i]
		result, _ := parseFrame(number)
		if expected != result {
			t.Errorf("Expected %v; but got %v", expected, result)
		}
	}
}

func Test_parse_miss_frame(t *testing.T) {
	numbers := []string{"1-", "5-", "--", "-"}
	expecteds := []int{1, 5, 0, 0}
	for i, number := range numbers {
		expected := expecteds[i]
		result, _ := parseFrame(number)
		if expected != result {
			t.Errorf("Expected %v; but got %v", expected, result)
		}
	}
}

func Test_spare_frame(t *testing.T) {
	numbers := []string{"1/", "5/", "-/", "9/"}
	for _, number := range numbers {
		expected := 10
		result, _ := parseFrame(number)
		if expected != result {
			t.Errorf("Expected %v; but got %v", expected, result)
		}
	}
}
