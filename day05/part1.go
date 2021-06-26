package day05

import (
	"io/ioutil"
	"strings"
)

func decodeSeat(spec string) int {
	id := 0
	for _, c := range spec {
		id = 2 * id
		switch c {
		case 'B', 'R':
			id++
		}
	}
	return id
}

func Part1() int {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	seatSpecs := strings.Split(string(content), "\n")
	maxSeatId := -1
	for _, spec := range seatSpecs {
		seatId := decodeSeat(spec)
		if seatId > maxSeatId {
			maxSeatId = seatId
		}
	}
	return maxSeatId
}
