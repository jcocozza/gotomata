package conway

import "github.com/jcocozza/gotomata/core"

func getCenter(width, height int) (int, int) {
	return width / 2, height / 2
}

func GliderConfig(width, height int) []core.Coordinate {
	conwayCenterR, conwayCenterC := getCenter(width, height)

	return []core.Coordinate{
		{conwayCenterR, conwayCenterC},
		{conwayCenterR + 1, conwayCenterC + 1},
		{conwayCenterR + 2, conwayCenterC - 1},
		{conwayCenterR + 2, conwayCenterC},
		{conwayCenterR + 2, conwayCenterC + 1},
	}

}

func AcornConfig(width, height int) []core.Coordinate {
	conwayCenterR, conwayCenterC := getCenter(width, height)

	return []core.Coordinate{
		{conwayCenterR - 1, conwayCenterC - 2},
		{conwayCenterR, conwayCenterC},
		{conwayCenterR + 1, conwayCenterC - 3},
		{conwayCenterR + 1, conwayCenterC - 2},
		{conwayCenterR + 1, conwayCenterC + 1},
		{conwayCenterR + 1, conwayCenterC + 2},
		{conwayCenterR + 1, conwayCenterC + 3},
	}
}
