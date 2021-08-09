package day11

import (
	"io/ioutil"
	"testing"
)

func TestPart2Example1(t *testing.T) {
	actual := solve(Part2{}, example1)
	expected := 26
	if actual != expected {
		t.Errorf("actual=%v expected=%v", actual, expected)
	}
}

func TestPart2(t *testing.T) {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	actual := solve(Part2{}, string(content))
	expected := 2059
	if actual != expected {
		t.Errorf("actual=%v expected=%v", actual, expected)
	}
}
