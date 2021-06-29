package day09

import "testing"

const testInput = `35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576`

func TestPart1Example(t *testing.T) {
	actual := solve(testInput, 5)
	expected := 127
	if actual != expected {
		t.Errorf("actual=%v expected=%v", actual, expected)
	}
}

func TestPart1(t *testing.T) {
	actual := Part1()
	expected := 1038347917
	if actual != expected {
		t.Errorf("actual=%v expected=%v", actual, expected)
	}
}
