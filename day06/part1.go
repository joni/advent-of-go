package day06

import (
	"io/ioutil"
	"strings"
	"unicode"
)

func solve(input string) int {
	groups := strings.Split(strings.TrimSpace(input), "\n\n")
	sumUniqueQs := 0
	for _, g := range groups {
		sumUniqueQs += countUniqueQuestions(g)
	}
	return sumUniqueQs
}

func countUniqueQuestions(group string) int {
	qMap := make(map[rune]bool)
	for _, q := range group {
		if !unicode.IsSpace(q) {
			qMap[q] = true
		}
	}
	return len(qMap)
}

func Part1() int {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return solve(string(content))
}
