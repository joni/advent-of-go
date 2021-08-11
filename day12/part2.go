package day12

import (
	"io/ioutil"
)

type Part2Ship struct {
	wpx  int
	wpy  int
	x, y int
}

func (ship Part2Ship) distance() int {
	return abs(ship.x) + abs(ship.y)
}

func (ship *Part2Ship) forward(n int) {
	ship.x += ship.wpx * n
	ship.y += ship.wpy * n
}

func (ship *Part2Ship) react(ins byte, n int) {
	switch ins {
	case 'N':
		ship.wpy += n
	case 'S':
		ship.wpy -= n
	case 'E':
		ship.wpx += n
	case 'W':
		ship.wpx -= n
	case 'L':
		ship.rotate(n)
	case 'R':
		ship.rotate(360 - n)

	}
}

func (ship *Part2Ship) rotate(angle int) {
	switch angle {
	case 90:
		ship.wpx, ship.wpy = -ship.wpy, ship.wpx
	case 180:
		ship.wpx, ship.wpy = -ship.wpx, -ship.wpy
	case 270:
		ship.wpx, ship.wpy = ship.wpy, -ship.wpx
	default:
		panic(angle)
	}
}

func Part2() int {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	ship := Part2Ship{wpx: 10, wpy: 1}
	return solve(string(content), &ship)
}
