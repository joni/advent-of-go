package day01

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func twoSum(numbers []int, target int) int {
	// This is a classic "2SUM" problem
	lookup := make(map[int]int, len(numbers))
	for i, n := range numbers {
		lookup[2020-n] = i
		if idx, ok := lookup[n]; ok {
			return numbers[idx] * n
		}
	}
	return 0
}

func Part1(filename string) int {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")
	numbers := make([]int, len(lines))
	for i, s := range lines {
		n, err := strconv.Atoi(s)
		if err == nil {
			numbers[i] = n
		}
	}

	return twoSum(numbers, 2020)
}
