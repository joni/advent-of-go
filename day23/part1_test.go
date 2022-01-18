package day23

import "testing"

func equal(left, right []int) bool {
	if len(left) != len(right) {
		return false
	}
	for idx := range left {
		if left[idx] != right[idx] {
			return false
		}
	}
	return true
}

func Test_Part1Example1(t *testing.T) {
	actual := solve([]int{3, 8, 9, 1, 2, 5, 4, 6, 7}, 10)
	expected := []int{9, 2, 6, 5, 8, 3, 7, 4}
	if !equal(actual, expected) {
		t.Errorf("actual %v != expected %v", actual, expected)
	}
}

func Test_Part1Example2(t *testing.T) {
	actual := solve([]int{3, 8, 9, 1, 2, 5, 4, 6, 7}, 100)
	expected := []int{6, 7, 3, 8, 4, 5, 2, 9}
	if !equal(actual, expected) {
		t.Errorf("actual %v != expected %v", actual, expected)
	}
}

func Test_Part1(t *testing.T) {
	actual := solve([]int{7, 3, 9, 8, 6, 2, 5, 4, 1}, 100)
	expected := []int{9, 4, 2, 3, 8, 6, 5, 7}
	if !equal(actual, expected) {
		t.Errorf("actual %v != expected %v", actual, expected)
	}
}
