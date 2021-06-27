package day07

import (
	"io/ioutil"
	"regexp"
	"strings"
)

type rule struct {
	bag         string
	contains    []string
	containedIn []string

	visited bool
}

func newRule(ruleStr string) *rule {
	twoWords := regexp.MustCompile(`(\w+ \w+) bags?`)
	bags := twoWords.FindAllStringSubmatch(ruleStr, -1)
	newRule := new(rule)
	for i, submatches := range bags {
		if i == 0 {
			newRule.bag = submatches[1]
		} else if submatches[1] != "no other" {
			newRule.contains = append(newRule.contains, submatches[1])
		}
	}
	return newRule
}

func fillContainedIn(ruleMap map[string]*rule) {
	for _, rule := range ruleMap {
		for _, containedBag := range rule.contains {
			ruleMap[containedBag].containedIn = append(ruleMap[containedBag].containedIn, rule.bag)
		}
	}
}

func dfs(ruleMap map[string]*rule, start string) int {
	queue := []string{start}
	ruleMap[start].visited = true
	visitCount := 0
	for {
		if len(queue) == 0 {
			return visitCount
		}
		current := queue[0]
		queue = queue[1:]

		for _, neigh := range ruleMap[current].containedIn {
			if !ruleMap[neigh].visited {
				visitCount++
				queue = append(queue, neigh)
				ruleMap[neigh].visited = true
			}
		}
	}
}

func solve(rules, bag string) int {
	rulesList := strings.Split(strings.TrimSpace(rules), "\n")
	ruleMap := make(map[string]*rule)
	for _, ruleStr := range rulesList {
		var currentRule = newRule(ruleStr)
		ruleMap[currentRule.bag] = currentRule
	}

	fillContainedIn(ruleMap)

	return dfs(ruleMap, bag)
}

func Part1() int {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return solve(string(content), "shiny gold")
}
