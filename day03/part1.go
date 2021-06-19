package day03

import (
	"io/ioutil"
	"strings"
)

func solve(input string, right int, down int) int {
	lines := strings.Split(input, "\n")
	xpos := 0
	treeCount := 0
	for ypos := 0; ypos < len(lines); ypos += down {
		line := lines[ypos]
		if line[xpos] == '#' {
			treeCount++
		}
		xpos = xpos + right
		if xpos >= len(line) {
			xpos = xpos - len(line)
		}
	}
	return treeCount
}

func Part1() int {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return solve(strings.TrimSpace(string(content)), 3, 1)
}
