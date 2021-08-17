package day15

import (
	"fmt"
)

func Part1(numbers []int) int {
	return solve(numbers, 2020)
}

func part1(numbers []int) int {
	for i := len(numbers); i < 2020; i++ {
		numbers = append(numbers, findPrev(numbers, numbers[i-1]))
	}
	fmt.Println(numbers)
	return numbers[2019]
}

func findPrev(numbers []int, n int) int {
	// fmt.Println("looking for", n, "in", numbers)
	for i := len(numbers) - 2; i >= 0; i-- {
		if n == numbers[i] {
			// fmt.Println("found at", i)
			return len(numbers) - 1 - i
		}
	}
	// fmt.Println("not found")
	return 0
}
