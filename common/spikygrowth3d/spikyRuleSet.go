package spikygrowth3d

import "github.com/jcocozza/gotomata/core"

func SpikyRuleSet(cell *core.Cell[int], neighbors []*core.Cell[int]) *core.Cell[int] {
	//0-3,7-9,11-13,18,21-22,24,26/13,17,20-26/4/M

	totalNeighbors := 0
	for _, neighbor := range neighbors {
		if neighbor.State == 3 {
			totalNeighbors++
		}
	}
	switch {
	case cell.State == 3 && ((totalNeighbors >= 0 && totalNeighbors <= 3) || (totalNeighbors >= 7 && totalNeighbors <= 9) || (totalNeighbors >= 11  && totalNeighbors <= 13) || totalNeighbors == 18 || totalNeighbors == 21 || totalNeighbors == 22 || totalNeighbors == 24 || totalNeighbors == 26):
		return &core.Cell[int]{State: 3, Coordinate: cell.Coordinate}
	case cell.State == 0 && (totalNeighbors == 4 || totalNeighbors == 13 || totalNeighbors == 17 || (totalNeighbors >= 20 && totalNeighbors <= 26)):
		return &core.Cell[int]{State: 3, Coordinate: cell.Coordinate}
	case cell.State == 0:
		return &core.Cell[int]{State: 0, Coordinate: cell.Coordinate}
	default:
		return &core.Cell[int]{State: cell.State - 1, Coordinate: cell.Coordinate}
	}
}
