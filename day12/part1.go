package day12

import (
	"io/ioutil"
	"strconv"
	"strings"
)

type Ship interface {
	react(ins byte, distance int)
	forward(distance int)
	distance() int
}

type Part1Ship struct {
	x       int
	y       int
	heading int
}

func abs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

func solve(input string, ship Ship) int {
	instructions := strings.Fields(input)
	for _, ins := range instructions {
		n, err := strconv.Atoi(ins[1:])
		if err != nil {
			panic(err)
		}
		if ins[0] == 'F' {
			ship.forward(n)
		} else {
			ship.react(ins[0], n)
		}

	}
	return ship.distance()
}

func (ship *Part1Ship) react(i byte, n int) {
	switch i {
	case 'N':
		ship.y += n
	case 'S':
		ship.y -= n
	case 'E':
		ship.x += n
	case 'W':
		ship.x -= n
	case 'L':
		ship.heading = (ship.heading + n) % 360
	case 'R':
		ship.heading = (360 + ship.heading - n) % 360
	}
}

func (ship *Part1Ship) forward(distance int) {
	switch ship.heading {
	case 0:
		ship.x += distance
	case 90:
		ship.y += distance
	case 180:
		ship.x -= distance
	case 270:
		ship.y -= distance
	}
}

func (ship Part1Ship) distance() int {
	return abs(ship.x) + abs(ship.y)
}

func Part1() int {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	ship := Part1Ship{}
	return solve(string(content), &ship)
}
