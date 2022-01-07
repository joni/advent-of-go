package day19

import (
	"bufio"
	"io/ioutil"
	"strconv"
	"strings"
)

type Rule interface {
	match(grammar *Grammar, message string) (int, bool)
}

type Alter []Concat

type Concat []int

type Literal string

type Rules [][]int

type Grammar map[int]Rule

func (alter Alter) match(grammar *Grammar, message string) (int, bool) {
	for _, alt := range alter {
		if end, ok := alt.match(grammar, message); ok {
			return end, true
		}
	}
	return 0, false
}

func (concat Concat) match(grammar *Grammar, message string) (int, bool) {
	allMatch := true
	end := 0
	for _, cc := range concat {
		consumed, ok := grammar.match(cc, message[end:])
		if !ok {
			allMatch = false
			break
		}
		end += consumed
	}
	if allMatch {
		return end, true
	} else {
		return 0, false
	}
}

func (literal Literal) match(grammar *Grammar, message string) (int, bool) {
	if strings.HasPrefix(message, string(literal)) {
		return len(literal), true
	} else {
		return 0, false
	}
}

func parseRhs(rhs string) Rule {
	if strings.HasPrefix(rhs, "\"") {
		return Literal(rhs[1 : len(rhs)-1])
	}
	rule := make(Alter, 0)
	altsSplit := strings.Split(rhs, " | ")
	for _, alt := range altsSplit {
		concat := make([]int, 0)
		concatSplit := strings.Split(alt, " ")
		for _, concatStr := range concatSplit {
			ruleNum, _ := strconv.Atoi(concatStr)
			concat = append(concat, ruleNum)
		}
		rule = append(rule, concat)
	}
	return rule
}

func parseRules(rulesStr string) Grammar {
	scanner := bufio.NewScanner(strings.NewReader(rulesStr))
	ruleMap := make(Grammar)
	for scanner.Scan() {
		lineStr := scanner.Text()
		lineSplit := strings.Split(lineStr, ": ")
		ruleNum, _ := strconv.Atoi(lineSplit[0])

		ruleMap[ruleNum] = parseRhs(lineSplit[1])
	}
	return ruleMap
}

func (rules *Grammar) match(ruleNum int, message string) (int, bool) {
	// fmt.Println(ruleNum, message)
	return (*rules)[ruleNum].match(rules, message)
}

func (rules *Grammar) matches(message string) bool {
	consumed, ok := rules.match(0, message)
	return ok && consumed == len(message)
	// fmt.Println(*rules, message)
	// return false
}

func solve1(input string) int {
	parts := strings.Split(input, "\n\n")
	rulesStr := parts[0]
	msgsStr := parts[1]
	messages := strings.Split(msgsStr, "\n")

	rules := parseRules(rulesStr)

	matchCount := 0
	for _, msg := range messages {
		if rules.matches(msg) {
			matchCount += 1
		}
	}

	return matchCount
}

func Part1() int {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return solve1(string(content))
}
