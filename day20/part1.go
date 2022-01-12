package day20

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const N = 10

type Tile struct {
	id      int
	content [][]byte
	found   bool
}

type Point struct {
	x int
	y int
}

func transpose(content [][]byte) {
	for i := range content {
		for j := 0; j < i; j++ {
			tmp := content[i][j]
			content[i][j] = content[j][i]
			content[j][i] = tmp
		}
	}
}

func reflect(content [][]byte) {
	n := len(content)
	for i := 0; i < n/2; i++ {
		tmp := content[i]
		content[i] = content[n-i-1]
		content[n-i-1] = tmp
	}
}

func (tile Tile) orient(pred func(Tile) bool) bool {
	if pred(tile) {
		return true
	}
	for i := 0; i < 4; i++ {
		transpose(tile.content)
		if pred(tile) {
			return true
		}
		reflect(tile.content)
		if pred(tile) {
			return true
		}
	}
	return false
}

func display(tile Tile) {
	fmt.Println("Tile", tile.id)
	for _, line := range tile.content {
		fmt.Println(string(line))
	}
}

func (top Tile) topMatchesBottom(bottom Tile) bool {
	for i := 0; i < N; i++ {
		if bottom.content[0][i] != top.content[N-1][i] {
			return false
		}
	}
	return true
}

func (left Tile) leftMatchesRight(right Tile) bool {
	for i := 0; i < N; i++ {
		if right.content[i][0] != left.content[i][N-1] {
			return false
		}
	}
	return true
}
func (bottom Tile) bottomMatchesTop(top Tile) bool {
	return top.topMatchesBottom(bottom)
}

func (right Tile) rightMatchesLeft(left Tile) bool {
	return left.leftMatchesRight(right)
}

func makeTile(tileStr string) Tile {
	lines := strings.Split(tileStr, "\n")
	id, _ := strconv.Atoi(lines[0][5:9])
	content := make([][]byte, 0, N)
	for _, line := range lines[1:] {
		content = append(content, []byte(line))
	}
	return Tile{id, content, false}
}

func readTiles(input string) []Tile {
	tileStrs := strings.Split(input, "\n\n")
	tiles := make([]Tile, 0, len(tileStrs))
	for _, tileStr := range tileStrs {
		tiles = append(tiles, makeTile(tileStr))
	}
	return tiles
}

func sortTiles(tiles []Tile) (map[Point]Tile, Point, Point) {

	tilesByCoords := make(map[Point]Tile)
	coords := make(map[int]Point)

	var dfs func(Tile)

	topLeft := Point{0, 0}
	bottomRight := Point{0, 0}

	dfs = func(tile Tile) {
		currentPoint := coords[tile.id]
		if currentPoint.x < topLeft.x {
			topLeft.x = currentPoint.x
		}
		if currentPoint.y < topLeft.y {
			topLeft.y = currentPoint.y
		}
		if currentPoint.x > bottomRight.x {
			bottomRight.x = currentPoint.x
		}
		if currentPoint.y > bottomRight.y {
			bottomRight.y = currentPoint.y
		}
		for _, candidate := range tiles {
			if _, ok := coords[candidate.id]; !ok {
				var pt Point
				if candidate.orient(tile.rightMatchesLeft) {
					pt = Point{currentPoint.x - 1, currentPoint.y}
				} else if candidate.orient(tile.leftMatchesRight) {
					pt = Point{currentPoint.x + 1, currentPoint.y}
				} else if candidate.orient(tile.topMatchesBottom) {
					pt = Point{currentPoint.x, currentPoint.y + 1}
				} else if candidate.orient(tile.bottomMatchesTop) {
					pt = Point{currentPoint.x, currentPoint.y - 1}
				} else {
					continue
				}
				tilesByCoords[pt] = candidate
				coords[candidate.id] = pt
				dfs(candidate)
			}
		}
	}

	tilesByCoords[Point{}] = tiles[0]
	coords[tiles[0].id] = Point{0, 0}
	dfs(tiles[0])

	return tilesByCoords, topLeft, bottomRight
}

func solve1(input string) int {
	tiles := readTiles(input)
	tilesByCoords, topLeft, bottomRight := sortTiles(tiles)
	topRight := Point{bottomRight.x, topLeft.y}
	bottomLeft := Point{topLeft.x, bottomRight.y}
	return tilesByCoords[topLeft].id * tilesByCoords[topRight].id * tilesByCoords[bottomLeft].id * tilesByCoords[bottomRight].id
}

func Part1() int {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return solve1(strings.TrimSpace(string(content)))
}
