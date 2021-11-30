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
	var cubes []Cube4
	for row, line := range lines {
		for col, char := range line {
			if char == '#' {
				cubes = append(cubes, Cube4{row, col, 0, 0})
			}
		}
	}
	for i := 0; i < 6; i++ {
		cubes = nextState2(cubes)
	}

	return len(cubes)
}

func nextState2(cubes []Cube4) []Cube4 {
	currCubes := make(map[Cube4]bool, len(cubes))
	cubeNeighs := make(map[Cube4]int, len(cubes)*20)
	for _, cube := range cubes {
		currCubes[cube] = true
		for _, neigh := range neighbors2(cube) {
			cubeNeighs[neigh] += 1
		}
	}
	var newState []Cube4
	for cube, count := range cubeNeighs {
		if currCubes[cube] {
			if count == 2 || count == 3 {
				newState = append(newState, cube)
			}
		} else if count == 3 {
			newState = append(newState, cube)
		}
	}
	return newState
}

func neighbors2(cube Cube4) []Cube4 {
	var neighs []Cube4
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
