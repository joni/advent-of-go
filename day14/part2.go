package day14

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"strings"
)

func Part2() int {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return solve2(string(content))
}

func solve2(input string) int {
	memory := make(map[int]int)
	var mask string
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		_, err := fmt.Sscanf(line, "mask = %s", &mask)
		if err == nil {
			continue
		}
		var addr, val int
		_, err = fmt.Sscanf(line, "mem[%d] = %d", &addr, &val)
		if err == nil {
			write(memory, mask, addr, val)
			continue
		}
	}

	sum := 0
	for _, cell := range memory {
		sum += cell
	}

	return sum
}

func write(memory map[int]int, mask string, address, val int) {
	addresses := genaddr(mask, address)
	for _, addr := range addresses {
		memory[addr] = val
	}
}

func genaddr(mask string, addr int) []int {
	addrs := []int{0}
	for idx, b := range mask {
		switch b {
		case '0':
			bit := (addr >> (35 - idx)) & 1
			for i, val := range addrs {
				addrs[i] = (val << 1) | bit
			}
		case '1':
			for i, val := range addrs {
				addrs[i] = (val << 1) | 1
			}
		case 'X':
			for i, val := range addrs {
				addrs[i] = (val << 1)
				addrs = append(addrs, (val<<1)|1)
			}
		}
	}

	return addrs
}
