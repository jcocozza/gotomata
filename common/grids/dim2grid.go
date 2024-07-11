package grids

import "github.com/jcocozza/gotomata/core"

func Dim2Grid[T comparable](width, height int, defaultState T, neighborsFunc core.GetNeighborsFunc) *core.Grid[T] {
	return &core.Grid[T]{
		BaseGrid: &core.BaseGrid[T]{
			Dimensions: []int{width, height},
			Cells: core.NewSparseCellGrid[T](),
			DefaultState: defaultState,
		},
		GetNeighborCoordinates: neighborsFunc,
	}
}

