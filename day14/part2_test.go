package day14

import "testing"

const example2 = `mask = 000000000000000000000000000000X1001X
mem[42] = 100
mask = 00000000000000000000000000000000X0XX
mem[26] = 1`

func TestPart2Example(t *testing.T) {
	actual := solve2(example2)
	expected := 208
	if actual != expected {
		t.Errorf("actual=%v expected=%v", actual, expected)
	}
}

func TestPart2(t *testing.T) {
	actual := Part2()
	expected := 4822600194774
	if actual != expected {
		t.Errorf("actual=%v expected=%v", actual, expected)
	}
}
