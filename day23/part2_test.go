package day23

import "testing"

func Test_Part2Example1(t *testing.T) {
	actual := solve2([]int{3, 8, 9, 1, 2, 5, 4, 6, 7})
	expected := 149245887792
	if actual != expected {
		t.Errorf("actual %v != expected %v", actual, expected)
	}
}

func Test_Part2(t *testing.T) {
	actual := solve2([]int{7, 3, 9, 8, 6, 2, 5, 4, 1})
	expected := 3072905352
	if actual != expected {
		t.Errorf("actual %v != expected %v", actual, expected)
	}
}
