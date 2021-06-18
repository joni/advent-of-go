package day02

import "testing"

func TestIsValid01(t *testing.T) {
	actual := isValid("1-3 a: abcde")
	expect := true
	if actual != expect {
		t.Errorf("actual %v, expected %v", actual, expect)
	}
}

func TestIsValid02(t *testing.T) {
	actual := isValid("1-3 b: cdefg")
	expect := false
	if actual != expect {
		t.Errorf("actual %v, expected %v", actual, expect)
	}
}

func TestPart1(t *testing.T) {
	actual := Part1("input.txt")
	expect := 628
	if actual != expect {
		t.Errorf("actual %v, expected %v", actual, expect)
	}
}
