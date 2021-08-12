package day13

import (
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func solve(input string) int {
	lines := strings.Split(input, "\n")
	arriveTime, err := strconv.Atoi(lines[0])
	if err != nil {
		panic(err)
	}

	shortestWait := math.MaxInt32
	shortestWaitBus := math.MaxInt32
	intervalsStr := strings.Split(lines[1], ",")

	for _, s := range intervalsStr {
		inter, err := strconv.Atoi(s)
		if err == nil {
			waitTime := (inter - arriveTime%inter) % inter
			if waitTime < shortestWait {
				shortestWait = waitTime
				shortestWaitBus = inter
			}
		}
	}

	return shortestWaitBus * shortestWait
}

func Part1() int {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return solve(string(content))
}
