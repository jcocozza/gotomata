// implements elementary cellular automata
package elementary

import "github.com/jcocozza/gotomata/core/base"

// implements base.RuleSet
type ElementaryRuleSet struct {
	ruleNumber uint8 
}

func newElementaryRuleSet(ruleNumber uint8) ElementaryRuleSet {
	if ruleNumber < 0 || ruleNumber > 255 {
		panic("Rule number must be between 0 and 255")
	}
	return ElementaryRuleSet{ruleNumber: ruleNumber}
}

func Btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

func (ers ElementaryRuleSet) GetRule(neighborhood []bool) base.Rule[bool] {
	left := Btoi(neighborhood[0])
	center := Btoi(neighborhood[1])
	right := Btoi(neighborhood[2])

    rule := func(nb []bool) bool {
        idx := left<<2 | center<<1 | right
        a := int(ers.ruleNumber >> idx & 1)
        return a == 1
    }
    return rule 
}

func elementaryNeighborhoods(cells [][]bool) base.Neighborhoods[bool] {
	return func(layer int) [][]bool {
		neighborhoods := make([][]bool, len(cells[layer]))
		for i := range cells[layer] {
			left := (i - 1 + len(cells[layer])) % len(cells[layer])
			right := (i + 1) % len(cells[layer])
			neighborhoods[i] = []bool{cells[layer][left], cells[layer][i], cells[layer][right]}
		}
		return neighborhoods
	}
}

func NewElementaryCellularAutomata(rule uint8, initLayer []bool, width, height int) *base.DiscreteCells {
	grid := base.Grid{
		Width:  width,
		Height: height,
	}

	rules := newElementaryRuleSet(rule)

	cells := &base.DiscreteCells{
		Grid:  grid,
		Rules: rules,
		Cells: [][]bool{},
		NFunc: nil,
	}
	cells.Init()
	nfunc := elementaryNeighborhoods(cells.Cells)
	cells.NFunc = nfunc
	return cells
}
