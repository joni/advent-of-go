package day03

import (
	"io/ioutil"
	"strings"
)

func solve2(input string) int {
	return solve(input, 1, 1) * solve(input, 3, 1) * solve(input, 5, 1) * solve(input, 7, 1) * solve(input, 1, 2)
}

func Part2() int {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return solve2(strings.TrimSpace(string(content)))
}
