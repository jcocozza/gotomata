package conway

import "github.com/jcocozza/gotomata/core"

func ConwayRuleSet(cell *core.Cell[bool], neighbors []*core.Cell[bool]) *core.Cell[bool] {
	totalLive := 0
	for _, nb := range neighbors {
		if nb.State {
			totalLive += 1
		}
	}

	switch {
	case totalLive < 2 && cell.State:
		return &core.Cell[bool]{State: false, Coordinate: cell.Coordinate}
    case (totalLive == 2 || totalLive == 3) && cell.State:
        return &core.Cell[bool]{State: true, Coordinate: cell.Coordinate}
    case totalLive > 3 && cell.State:
        return &core.Cell[bool]{State: false, Coordinate: cell.Coordinate}
    case !cell.State && totalLive == 3:
        return &core.Cell[bool]{State: true, Coordinate: cell.Coordinate}
    default:
        return &core.Cell[bool]{State: false, Coordinate: cell.Coordinate}
	}

}
