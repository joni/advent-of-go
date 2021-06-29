package day08

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type instruction struct {
	operator string
	argument int
}

func parseInstruction(line string) instruction {
	var ins instruction
	fmt.Sscan(line, &ins.operator, &ins.argument)
	return ins
}

type vm struct {
	pc  int // program counter
	acc int // accumulator
}

func (vm vm) run(instructions []instruction) (int, bool) {
	executedLines := make(map[int]bool)
	for {
		if vm.pc >= len(instructions) {
			return vm.acc, true
		}
		if _, found := executedLines[vm.pc]; found {
			return vm.acc, false
		}
		executedLines[vm.pc] = true
		ins := instructions[vm.pc]
		// fmt.Println(vm, ins)
		switch ins.operator {
		case "nop":
			vm.pc++
		case "acc":
			vm.acc += ins.argument
			vm.pc++
		case "jmp":
			vm.pc += ins.argument
		}
	}
}

func parseInstructions(input string) []instruction {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	instructions := make([]instruction, len(lines))
	for i, l := range lines {
		instructions[i] = parseInstruction(l)
	}
	return instructions
}

func solve(input string) int {
	instructions := parseInstructions(input)
	machine := vm{}
	acc, _ := machine.run(instructions)
	return acc
}

func Part1() int {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return solve(string(content))
}
