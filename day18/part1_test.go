package day18

import "testing"

const example = `1 + 2 * 3 + 4 * 5 + 6
2 * 3 + (4 * 5)
5 + (8 * 3 + 9 + 3 * 4 * 3)
5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))
((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2`

func TestPart1Example(t *testing.T) {
	actual := solve(Part1{}, example)
	expected := 71 + 26 + 437 + 12240 + 13632
	if actual != expected {
		t.Errorf("actual=%v expected=%v", actual, expected)
	}
}

func TestCalc1(t *testing.T) {
	actual := calc(Part1{}, "1 + 2 * 3 + 4 * 5 + 6")
	expected := 71
	if actual != expected {
		t.Errorf("actual=%v expected=%v", actual, expected)
	}
}

func TestCalc2(t *testing.T) {
	actual := calc(Part1{}, "2 * 3 + (4 * 5)")
	expected := 26
	if actual != expected {
		t.Errorf("actual=%v expected=%v", actual, expected)
	}
}

func TestPart1(t *testing.T) {
	actual := Solve(Part1{})
	expected := 3348222486398
	if actual != expected {
		t.Errorf("actual=%v expected=%v", actual, expected)
	}
}
