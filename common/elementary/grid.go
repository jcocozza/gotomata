package elementary

import "github.com/jcocozza/gotomata/core"

func elementaryNeighbors(length int) core.GetNeighborsFunc {
	nf := func(coord core.Coordinate) []core.Coordinate {
		left := core.Coordinate{(coord[0] - 1 + length) % length}
		right := core.Coordinate{(coord[0] + 1) % length}
		return []core.Coordinate{left, right}
	}
	return nf
}

func Dim1Grid[T comparable](length int, defaultState T) *core.Grid[T] {
	return &core.Grid[T]{
		BaseGrid: &core.BaseGrid[T]{
			Dimensions:   []int{length},
			Cells:        core.NewSparseCellGrid[T](),
			DefaultState: defaultState,
		},
		GetNeighborCoordinates: elementaryNeighbors(length),
	}
}
