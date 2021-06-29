package day08

import "io/ioutil"

func solve2(input string) int {
	instructions := parseInstructions(input)
	for i := range instructions {
		originalOperator := instructions[i].operator
		switch originalOperator {
		case "jmp":
			instructions[i].operator = "nop"
		case "nop":
			instructions[i].operator = "jmp"
		default:
			continue
		}
		machine := vm{}
		acc, ok := machine.run(instructions)
		if ok {
			return acc
		}
		instructions[i].operator = originalOperator
	}
	return -1
}

func Part2() int {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return solve2(string(content))
}
