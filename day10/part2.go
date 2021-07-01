package day10

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func countArrangementsOf1(ones int) int64 {
	// TODO figure out a generic solution
	switch ones {
	case 0, 1:
		return 1
	case 2:
		return 2
	case 3:
		return 4
	case 4:
		return 7
	default:
		return 0
	}
}

func part2(adapters []int) int64 {
	sort.Ints(adapters)
	diffs := make([]int, len(adapters))
	prev := 0
	var arrangements int64 = 1
	ones := 0
	for i, curr := range adapters {
		diffs[i] = curr - prev
		if diffs[i] == 3 {
			arrangements *= countArrangementsOf1(ones)
			ones = 0
		} else {
			ones++
		}
		prev = curr
	}
	fmt.Println(diffs)
	arrangements *= countArrangementsOf1(ones)
	return arrangements
}

func Part2() int64 {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	fields := strings.Fields(string(content))
	numbers := make([]int, len(fields))
	for i, f := range fields {
		numbers[i], _ = strconv.Atoi(f)
	}
	return part2(numbers)
}
