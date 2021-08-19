package day16

import "testing"

const example = `class: 1-3 or 5-7
row: 6-11 or 33-44
seat: 13-40 or 45-50

your ticket:
7,1,14

nearby tickets:
7,3,47
40,4,50
55,2,20
38,6,12`

func TestPart1Example(t *testing.T) {
	actual := solve1(example)
	expected := 71
	if actual != expected {
		t.Errorf("actual=%v expected=%v", actual, expected)
	}
}

func TestPart1(t *testing.T) {
	actual := Part1()
	expected := 22073
	if actual != expected {
		t.Errorf("actual=%v expected=%v", actual, expected)
	}
}
