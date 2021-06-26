package day05

import (
	"io/ioutil"
	"sort"
	"strings"
)

func Part2() int {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	seatSpecs := strings.Split(strings.TrimSpace(string(content)), "\n")
	seatIds := make([]int, len(seatSpecs))
	for i, spec := range seatSpecs {
		seatIds[i] = decodeSeat(spec)
	}
	sort.Ints(seatIds)
	prevSeatId := -1
	for i, seatId := range seatIds {
		if i != 0 && seatId != prevSeatId+1 {
			break
		}
		prevSeatId = seatId
	}
	return prevSeatId + 1
}
