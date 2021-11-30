package day17

import "testing"

const example = `.#.
..#
###`

func TestPart1Example(t *testing.T) {
	actual := solve1(example)
	expected := 112
	if actual != expected {
		t.Errorf("actual=%v expected=%v", actual, expected)
	}
}

func TestPart1(t *testing.T) {
	actual := Part1()
	expected := 359
	if actual != expected {
		t.Errorf("actual=%v expected=%v", actual, expected)
	}
}
