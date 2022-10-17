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

func Test_should_give_10_points_on_strike(t *testing.T) {
	number := "X|--|X|--"
	expected := 20
	result := parse(number)
	if expected != result {
		t.Errorf("Expected %v; but got %v", expected, result)
	}
}

func Test_should_give_10_points_on_spare(t *testing.T) {
	number := "4/|--|7/|--"
	expected := 20
	result := parse(number)
	if expected != result {
		t.Errorf("Expected %v; but got %v", expected, result)
	}
}

func Test_should_give_10_points_on_spare_and_double_next_throw(t *testing.T) {
	number := "4/|2-|7/|X"
	expected := 44
	result := parse(number)
	if expected != result {
		t.Errorf("Expected %v; but got %v", expected, result)
	}
}

func Test_should_give_10_points_on_strike_and_double_next_two_throw(t *testing.T) {
	number := "X|2-|X|33"
	expected := 36
	result := parse(number)
	if expected != result {
		t.Errorf("Expected %v; but got %v", expected, result)
	}
}
