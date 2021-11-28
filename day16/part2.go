package day16

import (
	"bufio"
	"errors"
	"io/ioutil"
	"regexp"
	"strings"
)

func Part2() int {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return solve2(string(content), 6)
}

func solve2(input string, limit int) int {
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

	scanner = bufio.NewScanner(strings.NewReader(paras[1]))
	scanner.Scan() // Skip "nearby tickets" line
	scanner.Scan()
	line := scanner.Text()
	yourTicket, _ := readTicket(fields, line)

	validTickets := [][]int{}

	scanner = bufio.NewScanner(strings.NewReader(paras[2]))
	scanner.Scan() // Skip "nearby tickets" line
	for scanner.Scan() {
		line := scanner.Text()
		ticket, err := readTicket(fields, line)
		if err == nil {
			validTickets = append(validTickets, ticket)
		}
	}

	validFields := make([][]bool, len(fields))
	for i := range validFields {
		validFields[i] = make([]bool, len(fields))
	}

	for fidx, field := range fields {
		for _, ticket := range validTickets {
			for vidx, value := range ticket {
				if !field.isValid(value) {
					validFields[fidx][vidx] = true
				}
			}
		}
	}

	fieldPermutation := make([]int, len(fields))
	loop := true
	for loop {
		loop = false
		for fidx, field := range validFields {
			countValid := 0
			validVIdx := -1
			for vidx, invalid := range field {
				if !invalid {
					countValid += 1
					validVIdx = vidx
				}
			}
			if countValid == 1 {
				loop = true
				fieldPermutation[fidx] = validVIdx
				for fidx := range validFields {
					validFields[fidx][validVIdx] = true
				}
			}
		}
	}

	ans := 1
	for _, vidx := range fieldPermutation[:limit] {
		ans *= yourTicket[vidx]
	}

	return ans
}

func readTicket(fields []Field, line string) ([]int, error) {
	valuesStr := strings.Split(line, ",")
	ticket := []int{}
	for _, valueStr := range valuesStr {
		value := s2i(valueStr)
		if !isValidForAnyField(fields, value) {
			return ticket, errors.New(valueStr)
		} else {
			ticket = append(ticket, value)
		}
	}
	return ticket, nil
}
