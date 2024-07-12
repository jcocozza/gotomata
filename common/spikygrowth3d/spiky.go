package spikygrowth3d

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jcocozza/gotomata/common/grids/dim3"
	"github.com/jcocozza/gotomata/core"
)

func Spiky(width, height, depth, steps int) *core.CellularAutomata[int] {
	grid := SpikyGrid(width, height, depth)
	return &core.CellularAutomata[int]{
		Grid:    grid,
		RuleSet: SpikyRuleSet,
		Steps:   steps,
	}
}

func ViewSpiky(ca *core.CellularAutomata[int]) {
	dim3.Visualizer(ca)
}

func stateToColor(state int) rl.Color {
	switch state {
	case 0:
		return rl.White
	case 1:
		return rl.LightGray
	case 2:
		return rl.Gray
	case 3:
		return rl.Black
	default:
		return rl.Blue
	}
}

