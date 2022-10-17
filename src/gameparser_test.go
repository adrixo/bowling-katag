package src

import "testing"

func Test_first(t *testing.T) {
	number := "23|42|32"
	expected := 16
	result := parse(number)
	if expected != result {
		t.Errorf("Expected %v; but got %v", expected, result)
	}
}
