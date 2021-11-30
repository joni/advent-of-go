package day17

import "testing"

func TestPart2Example(t *testing.T) {
	actual := solve2(example)
	expected := 848
	if actual != expected {
		t.Errorf("actual=%v expected=%v", actual, expected)
	}
}

func TestPart2(t *testing.T) {
	actual := Part2()
	expected := -1
	if actual != expected {
		t.Errorf("actual=%v expected=%v", actual, expected)
	}
}
