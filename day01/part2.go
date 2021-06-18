package day01

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func threeSum(numbers []int, target int) int {
	// This is a classic "3SUM" problem
	lookup := make(map[int]int, len(numbers))
	for i, n := range numbers {
		lookup[n] = i
		for j := 0; j < i; j++ {
			m := numbers[j]
			if idx, ok := lookup[2020-n-m]; ok {
				return numbers[idx] * n * m
			}
		}
	}
	return 0
}

func Part2(filename string) int {
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

	return threeSum(numbers, 2020)
}
