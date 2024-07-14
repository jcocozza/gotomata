package elementary

import "github.com/jcocozza/gotomata/core"

func SetCenterConfig(length int) []core.Coordinate {
	return []core.Coordinate{{length/2}}
}

func MainElementary(rule uint8, length, steps int, initConfig []core.Coordinate) {
	eca := ElementaryCellularAutomata(rule, length, steps)
	for _, coord := range initConfig {
		eca.Grid.SetCell(true, coord)
	}
	for i := 0; i < steps; i ++ {
		eca.Stepp()
		PrintECA(eca)
	}
}
