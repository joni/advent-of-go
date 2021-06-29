package day09

import "testing"

func TestPart2Example(t *testing.T) {
	actual := solve2(testInput, 127)
	expected := 62
	if actual != expected {
		t.Errorf("actual=%v expected=%v", actual, expected)
	}
}

func TestPart2(t *testing.T) {
	actual := Part2()
	expected := 137394018
	if actual != expected {
		t.Errorf("actual=%v expected=%v", actual, expected)
	}
}
