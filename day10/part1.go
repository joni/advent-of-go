package day10

import (
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func part1(adapters []int) int {
	sort.Ints(adapters)
	diff1 := 0
	diff3 := 1
	prev := 0
	for _, current := range adapters {
		switch current - prev {
		case 1:
			diff1++
		case 3:
			diff3++
		}
		prev = current
	}
	println(diff1, diff3)
	return diff1 * diff3
}

func Part1() int {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	fields := strings.Fields(string(content))
	numbers := make([]int, len(fields))
	for i, f := range fields {
		numbers[i], _ = strconv.Atoi(f)
	}
	return part1(numbers)
}
