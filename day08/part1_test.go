package day08

import "testing"

const exampleInput = `nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6
`

func TestSolve(t *testing.T) {
	actual := solve(exampleInput)
	expected := 5
	if actual != expected {
		t.Errorf("actual=%v expected=%v", actual, expected)
	}
}

func TestPart1(t *testing.T) {
	actual := Part1()
	expected := 2003
	if actual != expected {
		t.Errorf("actual=%v expected=%v", actual, expected)
	}
}
