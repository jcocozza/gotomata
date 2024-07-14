package crystals

import (
//	"fmt"

	"github.com/jcocozza/gotomata/common/grids"
	"github.com/jcocozza/gotomata/core"
)

func crystalNeighbors(width, height int) core.GetNeighborsFunc {
	offsets := [][]int{
		{1, 0}, {1, -1}, {0, -1},
		{-1, 0}, {-1, 1}, {0, 1},
	}

	crystalNeighborsFunc := func(coord core.Coordinate) []core.Coordinate {
		neighborhood := make([]core.Coordinate, 6)
		for i, offset := range offsets {
			q, r := coord[0], coord[1]
			nq, nr := q+offset[0], r+offset[1]
			neighborhood[i] = core.Coordinate{nq, nr}
//			fmt.Printf("Coord: %v :: Offset: %v -> Neighbor: (%d, %d)\n", coord, offset, nq, nr)
		}
		return neighborhood
	}

	return crystalNeighborsFunc
}

/*
func crystalNeighbors(width, height int) core.GetNeighborsFunc {
	offsets := [][]int{
		{1, 0}, {1, -1}, {0, -1},
		{-1, 0}, {-1, 1}, {0, 1},
	}

	crystalNeighborsFunc := func(coord core.Coordinate) []core.Coordinate {
		neighborhood := make([]core.Coordinate, 6)
		/*
			for v, offset := range offsets {
				q, r := coord[0], coord[1]
				nq, nr := q+offset[0], r+offset[1]
				neighborhood[v] = core.Coordinate{nq, nr}
			}
		for i := 0; i < 6; i++ {
			q, r := coord[0], coord[1]
			nq, nr := q+offsets[i][0], r+offsets[i][1]
			neighborhood[i] = core.Coordinate{nq, nr}
		}

		//fmt.Printf("Neighhorhood has %d members: %v\n", len(neighborhood), neighborhood)
		return neighborhood
	}

	return crystalNeighborsFunc
}
*/

func CrystalGrid(width, height int) *core.Grid[bool] {
	return grids.Dim2Grid(width, height, false, crystalNeighbors(width, height))
}
