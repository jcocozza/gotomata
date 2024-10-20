package conway

import (
	"github.com/jcocozza/gotomata/common/grids"
	"github.com/jcocozza/gotomata/core"
)

func conwayNeighbors(width, height int) core.GetNeighborsFunc {
	directions := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1} /*{0, 0}*/, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}
	conwayNeighborsFunc := func(coord core.Coordinate) []core.Coordinate {
		neighborhood := make([]core.Coordinate, 8)
		for v, dir := range directions {
			row, col := coord[0], coord[1]
			nr, nc := row+dir[0], col+dir[1]

			if nr >= 0 && nr < height && nc >= 0 && nc < width {
				neighborhood[v] = core.Coordinate{nr, nc}
			} else {
				neighborhood[v] = coord
			}
		}
		return neighborhood
	}
	return conwayNeighborsFunc
}

func dim2Grid[T comparable](width, height int, defaultState T, neighborFunc core.GetNeighborsFunc) *core.Grid[T] {
	return &core.Grid[T]{
		BaseGrid: &core.BaseGrid[T]{
			Dimensions:   []int{width, height},
			Cells:        core.NewSparseCellGrid[T](),
			DefaultState: defaultState,
		},
		GetNeighborCoordinates: neighborFunc,
	}
}

func ConwayGrid(width, height int) *core.Grid[bool] {
	return grids.Dim2Grid(width, height, false, conwayNeighbors(width, height))
	//return dim2Grid(width, height, false, conwayNeighbors(width, height))
}
