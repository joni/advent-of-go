package day24

import (
	"bufio"
	"io/ioutil"
	"strings"
)

type Tile struct {
	x, y int
}

type TileSet map[Tile]struct{}

func initial(input string) TileSet {
	scanner := bufio.NewScanner(strings.NewReader(input))
	tiles := make(TileSet)
	for scanner.Scan() {
		line := scanner.Text()
		tile := Tile{}
		for i := 0; i < len(line); i++ {
			c := line[i]
			if c == 'e' {
				tile.x++
			} else if c == 'w' {
				tile.x--
			} else {
				i++
				d := line[i]
				if c == 'n' {
					if d == 'e' {
						// ne
						tile.y++
					} else {
						// nw
						tile.x--
						tile.y++
					}
				} else {
					if d == 'e' {
						// se
						tile.x++
						tile.y--
					} else {
						// sw
						tile.y--
					}
				}
			}
		}
		if _, ok := tiles[tile]; ok {
			delete(tiles, tile)
		} else {
			tiles[tile] = struct{}{}
		}
	}
	return tiles
}

func solve(input string) int {
	tiles := initial(input)
	return len(tiles)
}

func Part1() int {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return solve(string(content))
}
