package day06

import (
	"io/ioutil"
	"strings"
)

func countCommonQuestions(group string) int {
	qMap := make(map[rune]bool)
	qs := strings.Split(group, "\n")
	for i, person := range qs {
		if i == 0 {
			for _, q := range person {
				qMap[q] = true
			}
		} else {
			for q := range qMap {
				if !strings.ContainsRune(person, q) {
					delete(qMap, q)
				}
			}
		}
	}
	return len(qMap)
}

func solve2(input string) int {
	groups := strings.Split(strings.TrimSpace(input), "\n\n")
	sumCommonQs := 0
	for _, g := range groups {
		sumCommonQs += countCommonQuestions(g)
	}
	return sumCommonQs
}

func Part2() int {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return solve2(string(content))
}
