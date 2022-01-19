package day24

import "io/ioutil"

func step(tiles TileSet) TileSet {
	newTiles := make(TileSet)
	count := make(map[Tile]int)
	for tile := range tiles {
		for _, neigh := range neighbors(tile) {
			count[neigh]++
		}
	}
	for tile, cnt := range count {
		if _, ok := tiles[tile]; ok {
			// black tile
			if cnt <= 2 {
				// keep black
				newTiles[tile] = struct{}{}
			}
		} else {
			// white tile
			if cnt == 2 {
				// flip to black
				newTiles[tile] = struct{}{}
			}
		}
	}
	return newTiles
}

func neighbors(tile Tile) []Tile {
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
