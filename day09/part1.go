package day09

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func isSumOfTwo(window []int, target int) bool {
	for i, n1 := range window {
		for _, n2 := range window[i:] {
			if n1+n2 == target {
				return true
			}
		}
	}
	return false
}

func solve(input string, prelen int) int {
	lines := strings.Fields(input)
	numbers := make([]int, len(lines))
	for i, n := range lines {
		numbers[i], _ = strconv.Atoi(n)
	}
	rest := numbers[prelen:]
	for i, n := range rest {
		window := numbers[i : i+prelen]
		if !isSumOfTwo(window, n) {
			return n
		}
		// fmt.Println(window, n)
	}
	return -1
}

func Part1() int {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return solve(string(content), 25)
}
