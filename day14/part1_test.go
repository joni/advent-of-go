package day14

import "testing"

const example = `mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0`

func TestPart1Example(t *testing.T) {
	actual := solve(example)
	expected := 165
	if actual != expected {
		t.Errorf("actual=%v expected=%v", actual, expected)
	}
}

func TestPart1(t *testing.T) {
	actual := Part1()
	expected := 11179633149677
	if actual != expected {
		t.Errorf("actual=%v expected=%v", actual, expected)
	}
}
