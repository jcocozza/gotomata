package dim3

import (
	"github.com/jcocozza/gotomata/core"
)

func Dim3CrystalConfig(width, height, depth int) []core.Coordinate {
	return []core.Coordinate{
		{width/2, height/2, depth/2},
		{width/2 + 1, height/2, depth/2},
		{width/2 + 2, height/2, depth/2},
	}	
}

func GeneralMain[T comparable](initConfig []core.Coordinate, aliveState int, ca *core.CellularAutomata[int]) {
	for _, coord := range initConfig {
		ca.Grid.SetCell(aliveState, coord)
	}
	VisualizeDim3(ca)
}
