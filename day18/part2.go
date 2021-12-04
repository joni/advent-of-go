package day18

import (
	"strconv"
	"strings"
)

type Part2 struct{}

func (part Part2) repl(expr string) string {
	return strconv.Itoa(part.subcalc(expr[1 : len(expr)-1]))
}

func (part Part2) subcalc(expr string) int {
	var operators []string
	var operands []int
	tokens := strings.Split(expr, " ")
	for _, t := range tokens {
		if t == "*" {
			for len(operators) > 0 && operators[len(operators)-1] == "+" {
				n := operands[len(operands)-1]
				m := operands[len(operands)-2]
				operands = operands[:len(operands)-2]
				operators = operators[:len(operators)-1]
				operands = append(operands, n+m)
			}
			operators = append(operators, t)
		} else if t == "+" {
			operators = append(operators, t)
		} else {
			operands = append(operands, atoi(t))
		}
	}

	for len(operators) > 0 {
		n := operands[len(operands)-1]
		m := operands[len(operands)-2]
		op := operators[len(operators)-1]
		operands = operands[:len(operands)-2]
		operators = operators[:len(operators)-1]
		if op == "+" {
			operands = append(operands, n+m)
		} else {
			operands = append(operands, n*m)

		}
	}

	return operands[0]
}
