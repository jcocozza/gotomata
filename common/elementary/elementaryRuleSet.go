package elementary

import "github.com/jcocozza/gotomata/core"

func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

func ElementaryRuleSet(ruleNumber uint8) core.RuleSet[bool] {
	ruleset := func(cell *core.Cell[bool], neighbors []*core.Cell[bool]) *core.Cell[bool] {
		left := btoi(neighbors[0].State)
		center := btoi(cell.State)
		right := btoi(neighbors[1].State)

		idx := left<<2 | center<<1 | right
		a := int(ruleNumber >> idx & 1)
		state := a == 1

		//       fmt.Println(left, center, right, "->", state)

		return &core.Cell[bool]{
			State: state,
			Coordinate: cell.Coordinate,
		}
	}
	return ruleset
}
