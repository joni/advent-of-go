package day02

import "testing"

func TestIsValid2_01(t *testing.T) {
	actual := isValid2("1-3 a: abcde")
	expect := true
	if actual != expect {
		t.Errorf("actual %v, expected %v", actual, expect)
	}
}

func TestIsValid2_02(t *testing.T) {
	actual := isValid2("1-3 b: cdefg")
	expect := false
	if actual != expect {
		t.Errorf("actual %v, expected %v", actual, expect)
	}
}

func TestIsValid2_03(t *testing.T) {
	actual := isValid2("2-9 c: ccccccccc")
	expect := false
	if actual != expect {
		t.Errorf("actual %v, expected %v", actual, expect)
	}
}

func TestPart2(t *testing.T) {
	actual := Part2("input.txt")
	expect := 705
	if actual != expect {
		t.Errorf("actual %v, expected %v", actual, expect)
	}
}
