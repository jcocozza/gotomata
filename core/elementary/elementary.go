package elementary

import (
//	"fmt"

	"github.com/jcocozza/gotomata/core/base"
)

type ElementaryCellularAutomata struct {
    RuleSet ElementaryRuleSet 
    Length int
    Grid *base.Grid[bool]
}

func NewECA(ruleNumber uint8, length int) *ElementaryCellularAutomata {
    return &ElementaryCellularAutomata{
        RuleSet: ElementaryRuleSet{ruleNumber: ruleNumber},
        Length: length,
        Grid: base.NewGrid[bool]([]int{length}),
    }
}

func NewECAFromGrid(ruleNumber uint8, length int, grid *base.Grid[bool]) *ElementaryCellularAutomata {
    return &ElementaryCellularAutomata{
        RuleSet: ElementaryRuleSet{ruleNumber: ruleNumber},
        Length: length,
        Grid: grid,
    }
}

func (eca *ElementaryCellularAutomata) GetNeighbors(cellLocation base.CellLocation) base.Neighborhood[bool] {
    // in this case, cell location corresponds exactly ot the place in the grid (b/c we have a 1d grid)
    center := eca.Grid.GetValue(cellLocation[0])
    left := eca.Grid.GetValue(  (cellLocation[0] - 1 + eca.Length) % eca.Length)
    right := eca.Grid.GetValue( (cellLocation[0]+ 1) % eca.Length)
    return base.Neighborhood[bool]{left, center, right}
}

func (eca *ElementaryCellularAutomata) GetRule(neighborhood base.Neighborhood[bool]) base.Rule[bool] {
    return eca.RuleSet.GetRule(neighborhood) 
}

func (eca *ElementaryCellularAutomata) Step() *ElementaryCellularAutomata {
    grid := base.NewGrid[bool](eca.Grid.Sizes)
    for i := range eca.Grid.Data {
        neighborhood := eca.GetNeighbors([]int{i})
        rule := eca.RuleSet.GetRule(neighborhood)
        result := rule(neighborhood)
        grid.SetValue(i, result)
    }    
    return NewECAFromGrid(eca.RuleSet.ruleNumber, eca.Length, grid) 
}
