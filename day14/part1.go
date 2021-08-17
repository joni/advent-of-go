package day14

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"strings"
)

func solve(input string) int {
	memory := make([]int, 1<<16)
	ones := 0
	zeros := 0
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		var mask string
		_, err := fmt.Sscanf(line, "mask = %s", &mask)
		if err == nil {
			ones, zeros = parseMask(mask)
			// println(mask, ones, zeros)
			continue
		}
		var addr, val int
		_, err = fmt.Sscanf(line, "mem[%d] = %d", &addr, &val)
		if err == nil {
			memory[addr] = (ones | val) & ^zeros
			// println(addr, val, memory[addr])
			continue
		}
	}

	sum := 0
	for _, cell := range memory {
		sum += cell
	}

	return sum
}

func parseMask(mask string) (ones, zeros int) {
	count := 0
	for i, c := range mask {
		idx := 35 - i
		switch c {
		case '1':
			ones = ones | (1 << idx)
		case '0':
			zeros = zeros | (1 << idx)
		default:
			count++
		}
	}
	println(count)
	return
}

func Part1() int {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return solve(string(content))
}
