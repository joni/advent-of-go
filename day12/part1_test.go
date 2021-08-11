package day12

import "testing"

const example1 = `F10
N3
F7
R90
F11`

func TestPart1Example(t *testing.T) {
	actual := solve(example1, &Part1Ship{})
	expected := 25
	if actual != expected {
		t.Errorf("actual=%v expected=%v", actual, expected)
	}
}

func TestPart1(t *testing.T) {
	actual := Part1()
	expected := 1482
	if actual != expected {
		t.Errorf("actual=%v expected=%v", actual, expected)
	}
}
