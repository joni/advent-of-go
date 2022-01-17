package day22

import "testing"

const infiniteExample = `Player 1:
43
19

Player 2:
2
29
14`

func Test_Part2Example(t *testing.T) {
	actual := solve2(example)
	expected := 291
	if actual != expected {
		t.Errorf("actual %v != expected %v", actual, expected)
	}
}

func Test_Part2InfiniteGame(t *testing.T) {
	actual := solve2(infiniteExample)
	expected := 2*43 + 19
	if actual != expected {
		t.Errorf("actual %v != expected %v", actual, expected)
	}
}

func Test_Part2(t *testing.T) {
	actual := Part2()
	expected := 32317
	if actual != expected {
		t.Errorf("actual %v != expected %v", actual, expected)
	}
}
