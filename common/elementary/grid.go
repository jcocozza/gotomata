package elementary

import (
	"github.com/jcocozza/gotomata/common/grids"
	"github.com/jcocozza/gotomata/core"
)

func elementaryNeighbors(length int) core.GetNeighborsFunc {
	nf := func(coord core.Coordinate) []core.Coordinate {
		left := core.Coordinate{(coord[0] - 1 + length) % length}
		right := core.Coordinate{(coord[0] + 1) % length}
		return []core.Coordinate{left, right}
	}
	return nf
}

func ElementaryGrid(length int) *core.Grid[bool] {
	return grids.Dim1Grid(length, false, elementaryNeighbors(length))
}
