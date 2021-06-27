package day07

import "testing"

const exampleInput = `light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.
`

func TestPart1Example1(t *testing.T) {
	actual := solve(exampleInput, "shiny gold")
	expected := 4
	if actual != expected {
		t.Errorf("actual = %v expected=%v", actual, expected)
	}
}

func TestPart1(t *testing.T) {
	actual := Part1()
	expected := 254
	if actual != expected {
		t.Errorf("actual = %v expected=%v", actual, expected)
	}
}
