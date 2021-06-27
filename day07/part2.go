package day07

import (
	"io/ioutil"
	"strings"
)

func dfs(ruleMap map[string]*rule, start string) int {
	bagCount := 1
	for neigh, count := range ruleMap[start].containsCount {
		bagCount += count * dfs(ruleMap, neigh)
	}
	return bagCount
}

func solve2(rules, bag string) int {
	rulesList := strings.Split(strings.TrimSpace(rules), "\n")
	ruleMap := make(map[string]*rule)
	for _, ruleStr := range rulesList {
		var currentRule = newRule(ruleStr)
		ruleMap[currentRule.bag] = currentRule
	}

	return dfs(ruleMap, bag) - 1
}

func Part2() int {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return solve2(string(content), "shiny gold")
}
