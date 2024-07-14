package elementary

import (
	"fmt"

	"github.com/jcocozza/gotomata/core"
)

func ElementaryCellularAutomata(rule uint8, length, steps int) *core.CellularAutomata[bool] {
	grid := ElementaryGrid(length)
	ruleset := ElementaryRuleSet(rule)

	return &core.CellularAutomata[bool]{
		Grid:    grid,
		RuleSet: ruleset,
		Steps:   steps,
	}
}

func PrintECA(eca *core.CellularAutomata[bool]) {
	s := ""
	coords := eca.Grid.AllCoordinates(nil)
	for _, coord := range coords {
		cell := eca.Grid.GetCell(coord)
		if cell.State {
			s += "█"
		} else {
			s += "░"
		}
	}
	fmt.Println(s)
}
