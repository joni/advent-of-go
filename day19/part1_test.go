package day19

import "testing"

const example = `0: 4 1 5
1: 2 3 | 3 2
2: 4 4 | 5 5
3: 4 5 | 5 4
4: "a"
5: "b"

ababbb
bababa
abbbab
aaabbb
aaaabbb
`

func TestPart1Example1(t *testing.T) {
	actual := solve1(example)
	expected := 2
	if actual != expected {
		t.Errorf("actual=%v expected=%v", actual, expected)
	}
}

func TestPart1(t *testing.T) {
	actual := Part1()
	expected := 220
	if actual != expected {
		t.Errorf("actual=%v expected=%v", actual, expected)
	}
}
