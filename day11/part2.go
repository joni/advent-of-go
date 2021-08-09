package day11

type Part2 struct {
}

func (part Part2) threshold() int {
	return 5
}

func (part Part2) countNeighbors(gen [][]byte, row, col int) int {
	count := 0
	directions := [][]int{
		{-1, -1}, {-1, 0}, {-1, +1},
		{0, -1} /*{0,0},*/, {0, +1},
		{+1, -1}, {+1, 0}, {+1, +1},
	}

	for _, dir := range directions {
		r, c := row+dir[0], col+dir[1]
		for r >= 0 && r < len(gen) && c >= 0 && c < len(gen[r]) && gen[r][c] == '.' {
			r, c = r+dir[0], c+dir[1]
		}
		if r >= 0 && r < len(gen) && c >= 0 && c < len(gen[r]) && gen[r][c] == '#' {
			count++
		}
	}

	return count
}
