package day25

import "testing"

func Test_Part1Example(t *testing.T) {
	actual := solve1(5764801, 17807724)
	expected := 14897079
	if actual != expected {
		t.Errorf("actual %v != expected %v", actual, expected)
	}
}

func Test_Part1(t *testing.T) {
	actual := solve1(11349501, 5107328)
	expected := 7936032
	if actual != expected {
		t.Errorf("actual %v != expected %v", actual, expected)
	}
}
