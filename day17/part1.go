package day17

import (
	"io/ioutil"
	"strings"
)

type Cube struct {
	x, y, z int
}

func solve1(input string) int {
	lines := strings.Split(input, "\n")
	var cubes []Cube
	for row, line := range lines {
		for col, char := range line {
			if char == '#' {
				cubes = append(cubes, Cube{row, col, 0})
			}
		}
	}
	for i := 0; i < 6; i++ {
		cubes = nextState(cubes)
	}

	return len(cubes)
}

func nextState(cubes []Cube) []Cube {
	currCubes := make(map[Cube]bool, len(cubes))
	cubeNeighs := make(map[Cube]int, len(cubes)*20)
	for _, cube := range cubes {
		currCubes[cube] = true
		for _, neigh := range neighbors(cube) {
			cubeNeighs[neigh] += 1
		}
	}
	var newState []Cube
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

func neighbors(cube Cube) []Cube {
	var neighs []Cube
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			for z := -1; z <= 1; z++ {
				if x != 0 || y != 0 || z != 0 {
					neighs = append(neighs, Cube{cube.x + x, cube.y + y, cube.z + z})
				}
			}
		}
	}
	return neighs
}

func Part1() int {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return solve1(string(content))
}
