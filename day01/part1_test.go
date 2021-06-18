package day01

import "testing"

func TestTwoSum(t *testing.T) {
	actual := twoSum([]int{1721, 979, 366, 299, 675, 1456}, 2020)
	expect := 514579
	if actual != expect {
		t.Errorf("actual %d, expected %d", actual, expect)
	}
}

func TestPart1(t *testing.T) {
	actual := Part1("input.txt")
	expect := 545379
	if actual != expect {
		t.Errorf("actual %d, expect %d", actual, expect)
	}
}
