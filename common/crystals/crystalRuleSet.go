package crystals

import "github.com/jcocozza/gotomata/core"

func CrystalRuleSet(cell *core.Cell[bool], neighbors []*core.Cell[bool]) *core.Cell[bool] {
	totalLive := 0
	for _, nb := range neighbors {
		if nb.State {
			totalLive += 1
		}
	}

	switch {
	// if the cell is dead and has exactly 1 neighbor, come to life
	// if the cell is alive, keep it alive
	case totalLive == 1 || cell.State:
		return &core.Cell[bool]{State: true, Coordinate: cell.Coordinate}
	default:
		return &core.Cell[bool]{State: false, Coordinate: cell.Coordinate}
	}
}
