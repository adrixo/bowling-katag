package src

import "testing"

func Test_should_sum_various_numeral_frames(t *testing.T) {
	number := "23|42|32"
	expected := 16
	result := parse(number)
	if expected != result {
		t.Errorf("Expected %v; but got %v", expected, result)
	}
}

func Test_should_ignore_miss_symbol(t *testing.T) {
	number := "--|4-|-2"
	expected := 6
	result := parse(number)
	if expected != result {
		t.Errorf("Expected %v; but got %v", expected, result)
	}
}
