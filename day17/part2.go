package day17

import (
	"io/ioutil"
	"strings"
)

type Cube4 struct {
	x, y, z, w int
}

func solve2(input string) int {
	lines := strings.Split(input, "\n")
	var cubes []Cube
	for row, line := range lines {
		for col, char := range line {
			if char == '#' {
				cubes = append(cubes, Cube4{row, col, 0, 0})
			}
		}
	}
	for i := 0; i < 6; i++ {
		cubes = nextState(cubes)
	}

	return len(cubes)
}

func (cube Cube4) neighbors() []Cube {
	var neighs []Cube
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			for z := -1; z <= 1; z++ {
				for w := -1; w <= 1; w++ {
					if x != 0 || y != 0 || z != 0 || w != 0 {
						neighs = append(neighs, Cube4{cube.x + x, cube.y + y, cube.z + z, cube.w + w})
					}
				}

			}
		}
	}
	return neighs
}

func Part2() int {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return solve2(string(content))
}
