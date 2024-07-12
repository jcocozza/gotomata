package crystal3d

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jcocozza/gotomata/common/grids/dim3"
	"github.com/jcocozza/gotomata/core"
)

func Crystal(width, height, depth, steps int) *core.CellularAutomata[bool] {
	grid := CrystalGrid(width, height, depth)
	return &core.CellularAutomata[bool]{
		Grid:    grid,
		RuleSet: CrystalGrowth,
		Steps:   steps,
	}
}

func ViewCrytal(ca *core.CellularAutomata[bool]) {
	dim3.Visualizer[bool](ca)
}

func distanceToColor(distance float64) rl.Color {

	//r := uint8((math.Sin(distance*0.2+0.5*math.Pi) + 1) * 127.5)
	//g := uint8((math.Sin(distance*0.2+1.5*math.Pi) + 1) * 127.5)
	//b := uint8((math.Sin(distance*0.2+2.5*math.Pi) + 1) * 127.5)

	r := uint8((math.Sin(distance*0.1) + 1) * 127.5)
	g := uint8((math.Sin(distance*0.1+2*math.Pi/3) + 1) * 127.5)
	b := uint8((math.Sin(distance*0.1+4*math.Pi/3) + 1) * 127.5)
	return rl.NewColor(r, g, b, 255)
}

