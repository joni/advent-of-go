package day18

import (
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type Calculator interface {
	subcalc(string) int
	repl(string) string
}

func solve(part Calculator, input string) int {
	lines := strings.Split(input, "\n")
	sum := 0
	for _, line := range lines {
		result := calc(part, line)
		sum += result
	}
	return sum
}

var parenRegex = regexp.MustCompile(`\([^()]*\)`)

func calc(part Calculator, expr string) int {
	for parenRegex.MatchString(expr) {
		expr = parenRegex.ReplaceAllStringFunc(expr, part.repl)
	}
	return part.subcalc(expr)
}

type Part1 struct{}

func (part Part1) repl(expr string) string {
	return strconv.Itoa(part.subcalc(expr[1 : len(expr)-1]))
}

func (part Part1) subcalc(expr string) int {
	tokens := strings.Split(expr, " ")
	acc := 0
	op := "+"
	for _, t := range tokens {
		if t == "+" || t == "*" {
			op = t
		} else {
			if op == "+" {
				acc += atoi(t)
			} else {
				acc *= atoi(t)
			}
		}
	}

	return acc
}

func atoi(a string) int {
	i, err := strconv.Atoi(a)
	if err != nil {
		panic(err)
	}
	return i
}

func Solve(part Calculator) int {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return solve(part, string(strings.TrimSpace(string(content))))
}
