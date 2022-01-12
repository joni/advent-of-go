package day20

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Image [][]byte

var monster = [][]byte{
	[]byte("                  # "),
	[]byte("#    ##    ##    ###"),
	[]byte(" #  #  #  #  #  #   "),
}

var monsterH = len(monster)
var monsterW = len(monster[0])

const monsterSize = 15

func isMonster(xOff int, yOff int, image Image) bool {
	c := byte('#')
	for y, line := range monster {
		for x, mc := range line {
			if mc == c && image[y+yOff][x+xOff] != c {
				return false
			}
		}
	}
	return true
}

func makeImage(w int, h int) Image {
	image := make(Image, 0, h)
	for i := 0; i < h; i++ {
		image = append(image, make([]byte, w))
	}
	return image
}

func displayImage(image Image) {
	for i, line := range image {
		fmt.Printf("%02d %s\n", i, string(line))
	}
}

func countMonsters(image Image) int {
	h := len(image)
	w := len(image[0])
	monsterCount := 0
	for r := 0; r < h-monsterH; r++ {
		for c := 0; c < w-monsterW; c++ {
			if isMonster(c, r, image) {
				monsterCount++
			}
		}
	}
	return monsterCount
}

func findMonsters(image Image) int {
	if count := countMonsters(image); count > 0 {
		return count
	}
	for i := 0; i < 4; i++ {
		transpose(image)
		if count := countMonsters(image); count > 0 {
			return count
		}

		reflect(image)
		if count := countMonsters(image); count > 0 {
			return count
		}
	}
	return 0
}

func countOctothorpe(image Image) int {
	count := 0
	for _, line := range image {
		for _, pixel := range line {
			if pixel == byte('#') {
				count++
			}
		}
	}
	return count
}

func createImageFromTiles(tiles []Tile) Image {
	sortedTiles, topLeft, bottomRight := sortTiles(tiles)
	w := (bottomRight.x - topLeft.x + 1) * (N - 2)
	h := (bottomRight.y - topLeft.y + 1) * (N - 2)
	completeImage := makeImage(w, h)

	for y := topLeft.y; y <= bottomRight.y; y++ {
		for x := topLeft.x; x <= bottomRight.x; x++ {
			tile := sortedTiles[Point{x, y}]
			tileOffX := (x - topLeft.x) * (N - 2)
			tileOffY := (y - topLeft.y) * (N - 2)
			for r := 0; r < N-2; r++ {
				copy(completeImage[tileOffY+r][tileOffX:tileOffX+N-2], tile.content[r+1][1:N-1])
			}
		}
	}
	return completeImage
}

func solve2(input string) int {
	tiles := readTiles(input)
	completeImage := createImageFromTiles(tiles)
	mosterCount := findMonsters(completeImage)
	waterChoppines := countOctothorpe(completeImage) - mosterCount*monsterSize
	return waterChoppines
}

func Part2() int {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return solve2(strings.TrimSpace(string(content)))
}
