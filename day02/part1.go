package day02

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func isValid(line string) bool {
	var min, max int
	var char rune
	var str string
	_, err := fmt.Sscanf(line, "%d-%d %c: %s", &min, &max, &char, &str)
	if err != nil {
		panic("pattern not matched")
	}
	count := 0
	for _, c := range str {
		if c == char {
			count++
		}
	}
	return min <= count && count <= max
}

func Part1(filename string) int {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(content)), "\n")
	count := 0
	for _, line := range lines {
		if isValid(line) {
			count++
		}
	}
	return count
}
