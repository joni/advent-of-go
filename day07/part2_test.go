package day07

import "testing"

func TestPart2Example1(t *testing.T) {
	actual := solve2(exampleInput, "shiny gold")
	expected := 32
	if actual != expected {
		t.Errorf("actual = %v expected=%v", actual, expected)
	}
}

const exampleInput2 = `shiny gold bags contain 2 dark red bags.
dark red bags contain 2 dark orange bags.
dark orange bags contain 2 dark yellow bags.
dark yellow bags contain 2 dark green bags.
dark green bags contain 2 dark blue bags.
dark blue bags contain 2 dark violet bags.
dark violet bags contain no other bags.
`

func TestPart2Example2(t *testing.T) {
	actual := solve2(exampleInput2, "shiny gold")
	expected := 126
	if actual != expected {
		t.Errorf("actual = %v expected=%v", actual, expected)
	}
}

func TestPart2(t *testing.T) {
	actual := Part2()
	expected := 6006
	if actual != expected {
		t.Errorf("actual = %v expected=%v", actual, expected)
	}
}
