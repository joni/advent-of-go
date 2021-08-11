package day12

import "testing"

func TestPart2Example(t *testing.T) {
	actual := solve(example1, &Part2Ship{wpx: 10, wpy: 1})
	expected := 286
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
