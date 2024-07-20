package continuous

import (
	"image"
	"image/color"
	"math"
	"math/rand"

	"github.com/jcocozza/gotomata/common/grids"
	"github.com/jcocozza/gotomata/core"
	"github.com/jcocozza/gotomata/utils"
)

func neighbors(length int) core.GetNeighborsFunc {
	nf := func(coord core.Coordinate) []core.Coordinate {
		left := core.Coordinate{(coord[0] - 1 + length) % length}
		right := core.Coordinate{(coord[0] + 1) % length}
		return []core.Coordinate{left, right}
	}
	return nf
}

var Neighbors = func(coord core.Coordinate) []core.Coordinate {
		left := core.Coordinate{(coord[0] - 1)}
		right := core.Coordinate{(coord[0] + 1)}
		return []core.Coordinate{left, right}
}

func grid(length int) *core.Grid[float64] {
	return grids.Dim1Grid[float64](length, 0, Neighbors)
}

var SimpleAverageRuleSet = func(cell *core.Cell[float64], neighbors []*core.Cell[float64]) *core.Cell[float64] {
	left := neighbors[0].State
	center := cell.State
	right := neighbors[1].State

	avg := (left + center + right) / 3

	return &core.Cell[float64]{
		State:      avg,
		Coordinate: cell.Coordinate,
	}
}

var KeepFractionalPartRuleSet = func(cell *core.Cell[float64], neighbors []*core.Cell[float64]) *core.Cell[float64] {
	left := neighbors[0].State
	center := cell.State
	right := neighbors[1].State

	avg := (left + center + right) / 3
	scaler := (float64(3)/float64(2)) * avg
	fracPart := scaler - math.Floor(scaler)

	return &core.Cell[float64]{
		State: fracPart,
		Coordinate: cell.Coordinate,
	}
}

func ContinuousCellularAutomata(length, steps int) *core.CellularAutomata[float64] {
	g := grid(length)
	return &core.CellularAutomata[float64]{
		Grid:    g,
		RuleSet: KeepFractionalPartRuleSet,
		Steps:   steps,
	}
}

func SetCenterConfig(length int) []core.Coordinate {
	return []core.Coordinate{{length / 2}}
}

func SetRandomConfig(length int) []core.Coordinate {
	initState := []core.Coordinate{}
	for i := 0; i < length; i++ {
		p :=  rand.Float64()
		if p > .5 {
			initState = append(initState, core.Coordinate{i})
		}
	}
	return initState
}

func MainContinuous(length, steps, scale int, initConfig []core.Coordinate) {
	ca := ContinuousCellularAutomata(length, steps)
	for _, coord := range initConfig {
		ca.Grid.SetCell(1, coord)
	}
	img := image.NewGray(image.Rect(0, 0, length*scale, steps*scale))
	AddContinuousToImage(ca, img, 0, scale, length)
	for i := 0; i < steps; i++ {
		ca.Stepp()
		if i < steps {
			AddContinuousToImage(ca, img, i+1, scale, length)
		}
	}
	utils.WritePNG(img, "_images/continuous.png")
}

func stateToColor(state float64) color.RGBA {
	gray := uint8((state) * 255) // 255 is the maximum value for 8-bit color channels
	return color.RGBA{R: gray, G: gray, B: gray, A: 255}
}

func AddContinuousToImage(ca *core.CellularAutomata[float64], img *image.Gray, stepNum, scale, length int) {
	coords := ca.Grid.AllCoordinates(nil)
	for i, coord := range coords {
		cell := ca.Grid.GetCell(coord)
		color := stateToColor(cell.State)
		for dx := 0; dx < scale; dx++ {
			for dy := 0; dy < scale; dy++ {
				img.Set((i*scale)+dx, (stepNum*scale) + dy, color)
			}
		}
	}
}
