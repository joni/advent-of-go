package day22

import (
	"testing"
)

const example = `Player 1:
9
2
6
3
1

Player 2:
5
8
4
7
10`

func Test_Part1Example(t *testing.T) {
	actual := solve(example)
	expected := 306
	if actual != expected {
		t.Errorf("actual %v != expected %v", actual, expected)
	}
}

func Test_Part1(t *testing.T) {
	actual := Part1()
	expected := 35202
	if actual != expected {
		t.Errorf("actual %v != expected %v", actual, expected)
	}
}
