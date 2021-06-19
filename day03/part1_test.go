package day03

import "testing"

func TestPart1Example(t *testing.T) {
	input := `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`
	actual := solve(input, 3, 1)
	expect := 7
	if actual != expect {
		t.Errorf("actual %v, expected %v", actual, expect)
	}
}

func TestPart1(t *testing.T) {
	actual := Part1()
	expect := 282
	if actual != expect {
		t.Errorf("actual %v, expected %v", actual, expect)
	}
}
