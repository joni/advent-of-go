package day13

import "testing"

const example = `939
7,13,x,x,59,x,31,19`

func TestPart1Example(t *testing.T) {
	actual := solve(example)
	expected := 295
	if actual != expected {
		t.Errorf("actual=%v expected=%v", actual, expected)
	}
}

func TestPart1(t *testing.T) {
	actual := Part1()
	expected := 4808
	if actual != expected {
		t.Errorf("actual=%v expected=%v", actual, expected)
	}
}
