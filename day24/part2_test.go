package day24

import "testing"

func Test_Part2Example(t *testing.T) {
	actual := solve2(example)
	expected := 2208
	if actual != expected {
		t.Errorf("actual %v != expected %v", actual, expected)
	}
}

func Test_Part2(t *testing.T) {
	actual := Part2()
	expected := 4353
	if actual != expected {
		t.Errorf("actual %v != expected %v", actual, expected)
	}
}
