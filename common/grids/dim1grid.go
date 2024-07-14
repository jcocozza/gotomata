package grids

import "github.com/jcocozza/gotomata/core"

func Dim1Grid[T comparable](length int, defaultState T, neighborsFunc core.GetNeighborsFunc) *core.Grid[T] {
	return &core.Grid[T]{
		BaseGrid: &core.BaseGrid[T]{
			Dimensions:   []int{length},
			Cells:        core.NewSparseCellGrid[T](),
			DefaultState: defaultState,
		},
		GetNeighborCoordinates: neighborsFunc,
	}
}
