package day01

import "testing"

func TestThreeSum(t *testing.T) {
	actual := threeSum([]int{1721, 979, 366, 299, 675, 1456}, 2020)
	expect := 241861950
	if actual != expect {
		t.Errorf("actual %d, expected %d", actual, expect)
	}
}

func TestPart2(t *testing.T) {
	actual := Part2("input.txt")
	expect := 257778836
	if actual != expect {
		t.Errorf("actual %d, expect %d", actual, expect)
	}
}
