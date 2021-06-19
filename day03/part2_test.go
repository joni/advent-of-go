package day03

import "testing"

func TestPart2Example(t *testing.T) {
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
	actual := solve2(input)
	expect := 336
	if actual != expect {
		t.Errorf("actual %v, expected %v", actual, expect)
	}
}

func TestPart2(t *testing.T) {
	actual := Part2()
	expect := 958815792
	if actual != expect {
		t.Errorf("actual %v, expected %v", actual, expect)
	}
}
