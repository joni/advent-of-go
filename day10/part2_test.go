package day10

import "testing"

func TestPart2Example1(t *testing.T) {
	actual := part2(example1)
	var expected int64 = 8
	if actual != expected {
		t.Errorf("actual=%v expected=%v", actual, expected)
	}
}

func TestPart2Example2(t *testing.T) {
	actual := part2(example2)
	var expected int64 = 19208
	if actual != expected {
		t.Errorf("actual=%v expected=%v", actual, expected)
	}
}

func TestPart2(t *testing.T) {
	actual := Part2()
	if actual != 12401793332096 {
		t.Errorf("actual=%v", actual)
	}
}
