package day06

import "testing"

func TestCountUnique(t *testing.T) {
	input := "ab\nac"
	actual := countUniqueQuestions(input)
	expected := 3
	if actual != expected {
		t.Errorf("actual=%v expected=%v", actual, expected)
	}
}

func TestPart1Example(t *testing.T) {
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
	expected := 11
	actual := solve(input)
	if actual != expected {
		t.Errorf("actual=%v expected=%v", actual, expected)
	}
}

func TestPart1(t *testing.T) {
	actual := Part1()
	expected := 6630

	if actual != expected {
		t.Errorf("actual=%v expected=%v", actual, expected)
	}
}
