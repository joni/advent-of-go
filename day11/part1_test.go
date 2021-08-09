package day11

import (
	"io/ioutil"
	"testing"
)

const example1 = `L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`

func TestPart1Example1(t *testing.T) {
	actual := solve(Part1{}, example1)
	expected := 37
	if actual != expected {
		t.Errorf("actual=%v expected=%v", actual, expected)
	}
}

func TestPart1(t *testing.T) {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	actual := solve(Part1{}, string(content))
	expected := 2289
	if actual != expected {
		t.Errorf("actual=%v expected=%v", actual, expected)
	}
}
