package day06

import "testing"

func TestCountCommon(t *testing.T) {
	input := "ab\nac"
	actual := countCommonQuestions(input)
	expected := 1
	if actual != expected {
		t.Errorf("actual=%v expected=%v", actual, expected)
	}
}

func TestPart2Example(t *testing.T) {
	input := `abc

a
b
c

ab
ac

a
a
a
a

b`
	expected := 6
	actual := solve2(input)
	if actual != expected {
		t.Errorf("actual=%v expected=%v", actual, expected)
	}
}

func TestPart2(t *testing.T) {
	actual := Part2()
	expected := 3437

	if actual != expected {
		t.Errorf("actual=%v expected=%v", actual, expected)
	}
}
