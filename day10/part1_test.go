package day10

import "testing"

var example1 []int = []int{16, 10, 15, 5, 1, 11, 7, 19, 6, 12, 4}

var example2 []int = []int{28, 33, 18, 42, 31, 14, 46, 20, 48, 47, 24, 23, 49, 45, 19, 38, 39, 11, 1, 32, 25, 35, 8, 17, 7, 9, 4, 2, 34, 10, 3}

func TestPart1Example1(t *testing.T) {
	actual := part1(example1)
	expected := 7 * 5
	if actual != expected {
		t.Errorf("actual=%v expected=%v", actual, expected)
	}
}

func TestPart1Example2(t *testing.T) {
	actual := part1(example2)
	expected := 22 * 10
	if actual != expected {
		t.Errorf("actual=%v expected=%v", actual, expected)
	}
}

func TestPart1(t *testing.T) {
	actual := Part1()
	expected := -1
	if actual != expected {
		t.Errorf("actual=%v expected=%v", actual, expected)
	}
}
