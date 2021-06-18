package day02

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func isValid2(line string) bool {
	var fst, snd int
	var char rune
	var str string
	_, err := fmt.Sscanf(line, "%d-%d %c: %s", &fst, &snd, &char, &str)
	if err != nil {
		return false
	}
	isFirstMatch := int(str[fst-1]) == int(char)
	isSecondMatch := int(str[snd-1]) == int(char)
	return isFirstMatch != isSecondMatch
}

func Part2(filename string) int {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(content)), "\n")
	count := 0
	for _, line := range lines {
		if isValid2(line) {
			count++
		}
	}
	return count
}
