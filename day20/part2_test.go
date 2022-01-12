package day20

import (
	"testing"
)

func TestPart2Example(t *testing.T) {
	actual := solve2(example1)
	expected := 273
	if actual != expected {
		t.Errorf("actual %v != expected %v", actual, expected)
	}
}
func TestPart2(t *testing.T) {
	actual := Part2()
	expected := 1649
	if actual != expected {
		t.Errorf("actual %v != expected %v", actual, expected)
	}
}
