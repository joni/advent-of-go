package day07

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

type rule struct {
	bag           string
	contains      []string
	containsCount map[string]int
	containedIn   []string

	visited bool
}

func newRule(ruleStr string) *rule {
	countColor := regexp.MustCompile(`(\d+ \w+ \w+) bags?`)
	bags := countColor.FindAllStringSubmatch(ruleStr, -1)
	newRule := new(rule)
	newRule.containsCount = make(map[string]int)
	var w1, w2 string
	fmt.Sscan(ruleStr, &w1, &w2)
	newRule.bag = w1 + " " + w2
	for _, submatches := range bags {
		var bagCount int
		fmt.Sscan(submatches[1], &bagCount, &w1, &w2)
		bagColor := w1 + " " + w2
		newRule.contains = append(newRule.contains, bagColor)
		newRule.containsCount[bagColor] = bagCount
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

func bfs(ruleMap map[string]*rule, start string) int {
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

	return bfs(ruleMap, bag)
}

func Part1() int {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return solve(string(content), "shiny gold")
}
