package conway

import "github.com/jcocozza/gotomata/core/base"

type GameOfLife struct {
	RuleSet ConwayRuleSet
	Width   int
	Height  int
	Grid    *base.Grid[bool]
}

func NewGOL(width, height int) *GameOfLife {
	return &GameOfLife{
		RuleSet: ConwayRuleSet{},
		Width:   width,
		Height:  height,
		Grid:    base.NewGrid[bool]([]int{width, height}),
	}
}
func NewGOLFromGrid(width, height int, grid *base.Grid[bool]) *GameOfLife {
	return &GameOfLife{
		RuleSet: ConwayRuleSet{},
		Width:   width,
		Height:  height,
		Grid:    grid,
	}
}

func (gol *GameOfLife) GetNeighbors(cellLocation base.CellLocation) base.Neighborhood[bool] {
	directions := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 0}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	neighborhood := make(base.Neighborhood[bool], 9)

	for v, dir := range directions {
		r, c := cellLocation[0], cellLocation[1]
		nr, nc := r+dir[0], c+dir[1]
		if nr >= 0 && nr < gol.Height && nc >= 0 && nc < gol.Width {
			neighborhood[v] = gol.Grid.GetValue(gol.Grid.Index([]int{nr, nc}))
		} else {
			neighborhood[v] = false
		}
	}
	return neighborhood
}

func (gol *GameOfLife) GetRule(neighborhood base.Neighborhood[bool]) base.Rule[bool] {
	return gol.RuleSet.GetRule(neighborhood)
}

func (gol *GameOfLife) Step() *GameOfLife {
	grid := base.NewGrid[bool]([]int{gol.Width, gol.Height})
	for i := range gol.Grid.Data {
		neighborhood := gol.GetNeighbors(gol.Grid.Coords(i))
		rule := gol.RuleSet.GetRule(neighborhood)
		result := rule(neighborhood)
		grid.SetValue(i, result)
	}
	return NewGOLFromGrid(gol.Width, gol.Height, grid)
}
/*
func (gol *GameOfLife) ToRows() [][]bool {
	rowLst := make([][]bool, gol.Height+1)

	counter := 0
	heightIdx := 0
	row := make([]bool, gol.Width)
	for _, cell := range gol.Grid.Data {
		row[counter] = cell
		if counter == gol.Width-1 {
			rowLst[heightIdx] = row
			counter = 0
			heightIdx += 1
			row = make([]bool, gol.Width)
		}
		counter += 1
	}
	return rowLst
}
*/

func (gol *GameOfLife) ToRows() [][]bool {
	rowLst := make([][]bool, gol.Height)

	for i := 0; i < gol.Height; i++ {
		row := make([]bool, gol.Width)
		for j := 0; j < gol.Width; j++ {
			row[j] = gol.Grid.GetValue(gol.Grid.Index([]int{i, j}))
		}
		rowLst[i] = row
	}
	return rowLst
}
