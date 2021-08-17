package day15

import (
	"testing"
)

func TestPart2Example(t *testing.T) {
	actual := Part2([]int{0, 3, 6})
	expected := 175594
	if actual != expected {
		t.Errorf("actual=%v expected=%v", actual, expected)
	}
}

func TestPart2Example2(t *testing.T) {
	actual := Part2([]int{1, 3, 2})
	expected := 2578
	if actual != expected {
		t.Errorf("actual=%v expected=%v", actual, expected)
	}
}

func TestPart2Example3(t *testing.T) {
	actual := Part2([]int{2, 1, 3})
	expected := 3544142
	if actual != expected {
		t.Errorf("actual=%v expected=%v", actual, expected)
	}
}

func TestPart2(t *testing.T) {
	actual := Part2([]int{16, 1, 0, 18, 12, 14, 19})
	expected := 16671510
	if actual != expected {
		t.Errorf("actual=%v expected=%v", actual, expected)
	}
}
