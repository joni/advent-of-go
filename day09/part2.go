package day09

import (
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func sumSmallestAndLargest(numbers []int) int {
	smallest := math.MaxInt32
	largest := math.MinInt32

	for _, n := range numbers {
		if n > largest {
			largest = n
		}
		if n < smallest {
			smallest = n
		}
	}
	return largest + smallest
}

func solve2(input string, target int) int {
	lines := strings.Fields(input)
	numbers := make([]int, len(lines))
	for i, n := range lines {
		numbers[i], _ = strconv.Atoi(n)
	}
	sum := numbers[0]
	left, right := 0, 1

	for {
		if sum == target {
			return sumSmallestAndLargest(numbers[left:right])
		}
		if sum < target {
			sum += numbers[right]
			right++
		} else {
			sum -= numbers[left]
			left++
		}
	}
}

func Part2() int {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return solve2(string(content), 1038347917)
}
