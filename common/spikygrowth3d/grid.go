package spikygrowth3d

import (
	"github.com/jcocozza/gotomata/common/grids"
	"github.com/jcocozza/gotomata/core"
)

func dim3NeighborsMoore(xLen, yLen, zLen int) core.GetNeighborsFunc {
	// 3d moore neighborhood
	directions := [][]int{
		{-1, -1, -1}, {-1, -1, 0}, {-1, -1, 1},
		{-1, 0, -1}, {-1, 0, 0}, {-1, 0, 1},
		{-1, 1, -1}, {-1, 1, 0}, {-1, 1, 1},
		{0, -1, -1}, {0, -1, 0}, {0, -1, 1},
		{0, 0, -1}, {0, 0, 1},
		{0, 1, -1}, {0, 1, 0}, {0, 1, 1},
		{1, -1, -1}, {1, -1, 0}, {1, -1, 1},
		{1, 0, -1}, {1, 0, 0}, {1, 0, 1},
		{1, 1, -1}, {1, 1, 0}, {1, 1, 1},
	}

	dim3NeighborsFunc := func(coord core.Coordinate) []core.Coordinate {
		neighborhood := make([]core.Coordinate, 26)
		for v, dir := range directions {
			x, y, z := coord[0], coord[1], coord[2]
			nx, ny, nz := x+dir[0], y+dir[1], z+dir[2]

			//if nx >= 0 && nx < xLen && ny >= 0 && ny < yLen && nz >= 0 && nz < zLen {
			neighborhood[v] = core.Coordinate{nx, ny, nz}
			//} else {
			//	neighborhood[v] = coord
			//}
		}
		return neighborhood
	}
	return dim3NeighborsFunc
}


func SpikyGrid(x, y, z int) *core.Grid[int] {
	return grids.Dim3Grid(x, y, z, 0, dim3NeighborsMoore(x, y, z))
}
