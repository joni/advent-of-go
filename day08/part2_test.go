package day08

import "testing"

func TestPart2Example(t *testing.T) {
	actual := solve2(exampleInput)
	expected := 8
	if actual != expected {
		t.Errorf("actual=%v expected=%v", actual, expected)
	}
}

func TestPart2(t *testing.T) {
	actual := Part2()
	expected := 1984
	if actual != expected {
		t.Errorf("actual=%v expected=%v", actual, expected)
	}
}
