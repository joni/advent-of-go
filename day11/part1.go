package day11

import (
	"fmt"
	"strings"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

type Part1 struct {
}

func (part Part1) threshold() int {
	return 4
}

func (part Part1) countNeighbors(gen [][]byte, row, col int) int {
	count := 0
	for r := max(0, row-1); r <= min(len(gen)-1, row+1); r++ {
		for c := max(0, col-1); c <= min(len(gen[r])-1, col+1); c++ {
			if !(r == row && c == col) && gen[r][c] == '#' {
				count++
			}
		}
	}
	return count
}

type WaitingRoom interface {
	countNeighbors(gen [][]byte, row, col int) int
	threshold() int
}

func solve(wr WaitingRoom, input string) int {
	lines := strings.Fields(input)
	gen1 := make([][]byte, len(lines))
	for i, line := range lines {
		gen1[i] = []byte(line)
	}

	oldgen := gen1

	for j := 1; j < 9999999; j++ {

		newgen := make([][]byte, len(oldgen))
		for r, line := range oldgen {
			newgen[r] = make([]byte, len(line))
			for c, seat := range line {
				nc := wr.countNeighbors(oldgen, r, c)
				if seat == 'L' && nc == 0 {
					newgen[r][c] = '#'
				} else if seat == '#' && nc >= wr.threshold() {
					newgen[r][c] = 'L'
				} else {
					newgen[r][c] = seat
				}
			}
		}

		fmt.Println("Generation", j)

		count := 0
		for r, line := range newgen {
			for c, seat := range line {
				if oldgen[r][c] != seat {
					count++
				}
			}
			fmt.Println(string(line))
		}
		fmt.Println()
		if count == 0 {
			break
		} else {
			oldgen = newgen
		}
	}

	count := 0
	for _, line := range oldgen {
		for _, seat := range line {
			if seat == '#' {
				count++
			}
		}
	}
	return count
}
