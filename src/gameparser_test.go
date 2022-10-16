package src

import "testing"

func Test_example(t *testing.T) {
	expected := 0
	result := parse("")
	if expected != result {
		t.Errorf("Expected %v; but got %v", expected, result)
	}
}
