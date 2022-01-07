package day19

import (
	"io/ioutil"
	"strings"
)

func (rules *Grammar) matchRepeated(ruleNum int, message string) (int, int) {
	idx := 0
	count := 0
	for {
		consumed, ok := rules.match(ruleNum, message[idx:])
		if !ok {
			return idx, count
		}
		count += 1
		idx += consumed
	}
}

func (rules *Grammar) matches2(message string) bool {
	consumed42, count42 := rules.matchRepeated(42, message)
	consumed31, count31 := rules.matchRepeated(31, message[consumed42:])
	return count42 > count31 && count31 > 0 && consumed42+consumed31 == len(message)
}

func solve2(input string) int {
	parts := strings.Split(input, "\n\n")
	rulesStr := parts[0]
	msgsStr := parts[1]
	messages := strings.Split(msgsStr, "\n")

	rules := parseRules(rulesStr)
	rules[8] = Alter{Concat{42}, Concat{42, 8}}
	rules[11] = Alter{Concat{42, 31}, Concat{42, 11, 31}}

	matchCount := 0
	for _, msg := range messages {
		if rules.matches2(msg) {
			matchCount += 1
		}
	}

	return matchCount
}

func Part2() int {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return solve2(string(content))
}
