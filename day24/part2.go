package day24

import "io/ioutil"

func step(tiles TileSet) TileSet {
	newTiles := NewTileSet()
	count := make(map[Tile]int)
	for tile := range tiles {
		for _, neigh := range tile.neighbors() {
			count[neigh]++
		}
	}
	for tile, cnt := range count {
		if tiles.Contains(tile) {
			// black tile
			if cnt <= 2 {
				// keep black
				newTiles.Add(tile)
			}
		} else {
			// white tile
			if cnt == 2 {
				// flip to black
				newTiles.Add(tile)
			}
		}
	}
	return newTiles
}

func (tile Tile) neighbors() []Tile {
	return []Tile{
		{tile.x + 1, tile.y},
		{tile.x - 1, tile.y},
		{tile.x, tile.y + 1},
		{tile.x, tile.y - 1},
		{tile.x - 1, tile.y + 1},
		{tile.x + 1, tile.y - 1},
	}
}

func solve2(input string) int {
	tiles := initial(input)

	for i := 0; i < 100; i++ {
		tiles = step(tiles)
	}

	return len(tiles)
}

func Part2() int {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return solve2(string(content))
}
