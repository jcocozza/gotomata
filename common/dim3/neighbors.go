package dim3

import (
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

func dim3NeighborsVN(xLen, yLen, zLen int) core.GetNeighborsFunc {
	// 3d VN neighborhood
	directions := [][]int{
		{-1, 0, 0}, {1, 0, 0}, // x direction
		{0, -1, 0}, {0, 1, 0}, // y direction
		{0, 0, -1}, {0, 0, 1}, // z direction
	}

	dim3NeighborsFunc := func(coord core.Coordinate) []core.Coordinate {
		neighborhood := make([]core.Coordinate, 6)
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
