package dim3

import (
	"github.com/jcocozza/gotomata/common/grids"
	"github.com/jcocozza/gotomata/common/grids/dim3viz"
	"github.com/jcocozza/gotomata/core"
)

func A(x,y,z int) *core.Grid[int] {
	return grids.Dim3Grid(x,y,z, 0, dim3NeighborsMoore(x,y,z))
}

func Ao(x,y,z int) *core.Grid[int] {
	return grids.Dim3Grid(x,y,z, 0, dim3NeighborsVN(x,y,z))
}


func Crystal(width, height, depth, steps int) *core.CellularAutomata[int] {
	grid := grids.Dim3Grid(width, height, depth, 0, dim3NeighborsVN(width, height, depth))
	return &core.CellularAutomata[int] {
		Grid: grid,
		RuleSet: GenerateRuleSet([]int{0,1,2,3,4,5,6}, []int{1,3}, 2),
		Steps: steps,
	}
}

func R445(width, height, depth, steps int) *core.CellularAutomata[int] {
	grid := grids.Dim3Grid(width, height, depth, 0, dim3NeighborsMoore(width, height, depth))
	return &core.CellularAutomata[int] {
		Grid: grid,
		RuleSet: GenerateRuleSet([]int{4}, []int{4}, 5),
		Steps: steps,
	}
}

func Amoeba(width, height, depth, steps int) *core.CellularAutomata[int] {
	grid := grids.Dim3Grid(width, height, depth, 0, dim3NeighborsMoore(width, height, depth))
	return &core.CellularAutomata[int] {
		Grid: grid,
		RuleSet: GenerateRuleSet([]int{9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26}, []int{5,6,7,12,13,15}, 5),
		Steps: steps,
	}
}

func R678678(width, height, depth, steps int) *core.CellularAutomata[int] {
	grid := grids.Dim3Grid(width, height, depth, 0, dim3NeighborsMoore(width, height, depth))
	return &core.CellularAutomata[int] {
		Grid: grid,
		RuleSet: GenerateRuleSet([]int{0,1,2,3,7,8,9,11,12,13,18,21,22,24,26}, []int{4,13,17,20,21,22,23,24,26}, 4),
		Steps: steps,
	}
}


func VisualizeDim3(ca *core.CellularAutomata[int]) {
	dim3viz.Visualizer(ca)
}
