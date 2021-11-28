package day16

import "testing"

const example2 = `class: 0-1 or 4-19
row: 0-5 or 8-19
seat: 0-13 or 16-19

your ticket:
11,12,13

nearby tickets:
3,9,18
15,1,5
5,14,9`

func TestPart2Example(t *testing.T) {
	actual := solve2(example2, 3)
	expected := 11 * 12 * 13
	if actual != expected {
		t.Errorf("actual=%v expected=%v", actual, expected)
	}
}

func TestPart2(t *testing.T) {
	actual := Part2()
	expected := 1346570764607
	if actual != expected {
		t.Errorf("actual=%v expected=%v", actual, expected)
	}
}
