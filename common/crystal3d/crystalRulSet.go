package crystal3d

import (

	"github.com/jcocozza/gotomata/core"
)

func AmeobaRuleSet(cell *core.Cell[int], neighbors []*core.Cell[int]) *core.Cell[int] {
	//9-26/5-7,12-13,15/5/M
	// cells with 9,10,...,26 neighbors survive
	// cells with 5,...,7 OR 12,13, OR 15 have a new cell born
	// cells have 5 total states 0 (dead) - 4 (alive)

	totalNeighbors := 0
	for _, neighbor := range neighbors {
		if neighbor.State == 4 {
			totalNeighbors++
		}
	}

	switch {
	case cell.State == 4:
		if totalNeighbors >= 9 && totalNeighbors <= 26 {// cell survives
			return &core.Cell[int]{State: 4, Coordinate: cell.Coordinate}
		} else { // decrement by 1
			return &core.Cell[int]{State: cell.State - 1, Coordinate: cell.Coordinate}
		}
	case cell.State == 0:
		if (totalNeighbors >= 5 && totalNeighbors <= 7) || (totalNeighbors == 12) || (totalNeighbors == 13) || ( totalNeighbors == 15) { // new cell is born
			return &core.Cell[int]{State: 4, Coordinate: cell.Coordinate}
		} else {
			return &core.Cell[int]{State: 0, Coordinate: cell.Coordinate}
		}
	default: // decrement by 1
		return &core.Cell[int]{State: cell.State - 1, Coordinate: cell.Coordinate}
	}
}

func R445RuleSet(cell *core.Cell[int], neighbors []*core.Cell[int]) *core.Cell[int] {
	// 4/4/5/M

	totalNeighbors := 0
	for _, neighbor := range neighbors {
		if neighbor.State == 1 {
			totalNeighbors++
		}
	}

	switch {
	case cell.State == 4:
		if totalNeighbors == 4 {
			return &core.Cell[int]{State: 4, Coordinate: cell.Coordinate}
		} else {
			return &core.Cell[int]{State: cell.State - 1, Coordinate: cell.Coordinate}
		}
	case cell.State == 0:
		if totalNeighbors == 4 {
			return &core.Cell[int]{State: 4, Coordinate: cell.Coordinate}
		} else {
			return &core.Cell[int]{State: 0, Coordinate: cell.Coordinate}
		}
	default:
		return &core.Cell[int]{State: cell.State - 1, Coordinate: cell.Coordinate}
	}
}

func CrystalGrowth(cell *core.Cell[bool], neighbors []*core.Cell[bool]) *core.Cell[bool] {
	// 0-6/1,3/2/VN

	totalNeighbors := 0
	for _, neighbor := range neighbors {
		if neighbor.State {
			totalNeighbors++
		}
	}
	switch {
	case totalNeighbors <= 6 && cell.State:
		return &core.Cell[bool]{State: true, Coordinate: cell.Coordinate}
	case (totalNeighbors == 1 || totalNeighbors == 3) && !cell.State:
		return &core.Cell[bool]{State: true, Coordinate: cell.Coordinate}
	default:
		return &core.Cell[bool]{State: false, Coordinate: cell.Coordinate}
	}
}
