package day16

import (
	"bufio"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func s2i(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

type Field struct {
	validationRules []Rule
}

type Rule struct {
	min int
	max int
}

func (rule Rule) isValid(value int) bool {
	if rule.min <= value && value <= rule.max {
		return true
	}
	return false
}

func (field Field) isValid(value int) bool {
	for _, rule := range field.validationRules {
		if rule.isValid(value) {
			return true
		}
	}
	return false
}

func isValidForAnyField(fields []Field, value int) bool {
	for _, f := range fields {
		if f.isValid(value) {
			return true
		}
	}
	return false
}

func solve1(input string) int {
	paras := strings.Split(input, "\n\n")
	rangeRe := regexp.MustCompile(`(\d+)-(\d+)`)

	fields := []Field{}

	scanner := bufio.NewScanner(strings.NewReader(paras[0]))
	for scanner.Scan() {
		line := scanner.Text()
		matches := rangeRe.FindAllStringSubmatch(line, -1)
		rules := make([]Rule, len(matches))
		for i, m := range matches {
			rules[i] = Rule{min: s2i(m[1]), max: s2i(m[2])}
		}

		f := Field{rules}
		fields = append(fields, f)

	}

	sumInvalid := 0
	scanner = bufio.NewScanner(strings.NewReader(paras[2]))
	scanner.Scan() // Skip "nearby tickets" line
	for scanner.Scan() {
		line := scanner.Text()
		valuesStr := strings.Split(line, ",")
		for _, valueStr := range valuesStr {
			value := s2i(valueStr)
			if !isValidForAnyField(fields, value) {
				sumInvalid += value
			}
		}
	}

	return sumInvalid
}

func Part1() int {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return solve1(string(content))
}
