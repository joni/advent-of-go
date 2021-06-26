package day05

import "testing"

func TestPart2(t *testing.T) {
	actual := Part2()
	expected := -1
	if actual != expected {
		t.Errorf("actual=%v expected=%v", actual, expected)
	}
}
