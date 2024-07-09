package randomwalk

import (
	"github.com/jcocozza/gotomata/common/grids"
	"github.com/jcocozza/gotomata/core"
)

func randomWalkNeighbors(width, height int) core.GetNeighborsFunc {
	// use von Neumann neighborhood
	/*
	directions := [][]int{
			{-1, 0},
		{0, -1}, {0, 1},
			{1, 0},
	}*/
	directions := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1} /*{0, 0}*/, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	randomWalkNeighborsFunc := func(coord core.Coordinate) []core.Coordinate {
		neighborhood := make([]core.Coordinate, 8)
		for v, dir := range directions {
			row, col := coord[0], coord[1]
			nr, nc := row + dir[0], col + dir[1]

			if nr >= 0 && nr < height && nc >= 0 && nc < width {
				neighborhood[v] = core.Coordinate{nr, nc}
			} else {
				neighborhood[v] = coord
			}
		}
		return neighborhood
	}
	return randomWalkNeighborsFunc
}

func randomWalkGrid(width, height int) *core.Grid[bool] {
	return grids.Dim2Grid(width, height, false, randomWalkNeighbors(width, height))
}
