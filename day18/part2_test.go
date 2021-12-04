package day18

import "testing"

func TestPart2Example(t *testing.T) {
	actual := solve(Part2{}, example)
	expected := 231 + 46 + 1445 + 669060 + 23340
	if actual != expected {
		t.Errorf("actual=%v expected=%v", actual, expected)
	}
}

func TestCalc3(t *testing.T) {
	actual := calc(Part2{}, "1 + 2 * 3 + 4 * 5 + 6")
	expected := 231
	if actual != expected {
		t.Errorf("actual=%v expected=%v", actual, expected)
	}
}

func TestPart2(t *testing.T) {
	actual := Solve(Part2{})
	expected := 43423343619505
	if actual != expected {
		t.Errorf("actual=%v expected=%v", actual, expected)
	}
}
