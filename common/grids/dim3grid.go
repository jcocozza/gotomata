package grids

import "github.com/jcocozza/gotomata/core"

func Dim3Grid[T comparable](x, y, z int, defaultState T, neighborsFunc core.GetNeighborsFunc) *core.Grid[T] {
	return &core.Grid[T]{
		BaseGrid: &core.BaseGrid[T]{
			Dimensions: []int{x, y, z},
			Cells: core.NewSparseCellGrid[T](),
			DefaultState: defaultState,
		},
		GetNeighborCoordinates: neighborsFunc,
	}
}
