package day15

import (
	"testing"
)

func TestPart1Example(t *testing.T) {
	actual := Part1([]int{0, 3, 6})
	expected := 436
	if actual != expected {
		t.Errorf("actual=%v expected=%v", actual, expected)
	}
}

func TestPart1Example2(t *testing.T) {
	actual := Part1([]int{1, 3, 2})
	expected := 1
	if actual != expected {
		t.Errorf("actual=%v expected=%v", actual, expected)
	}
}

func TestPart1Example3(t *testing.T) {
	actual := Part1([]int{2, 1, 3})
	expected := 10
	if actual != expected {
		t.Errorf("actual=%v expected=%v", actual, expected)
	}
}

func TestPart1(t *testing.T) {
	actual := Part1([]int{16, 1, 0, 18, 12, 14, 19})
	expected := 929
	if actual != expected {
		t.Errorf("actual=%v expected=%v", actual, expected)
	}
}
