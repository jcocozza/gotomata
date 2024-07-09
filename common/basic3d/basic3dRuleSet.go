package basic3d

import "github.com/jcocozza/gotomata/core"

func Basic3dRuleSet(cell *core.Cell[bool], neighbors []*core.Cell[bool]) *core.Cell[bool] {
	first4Alive := 0
	for i := 0; i < 4; i++ {
		if neighbors[i].State {
			first4Alive += 1
		}
	}

	second4Alive := 0
	for i := 4; i < 8; i++ {
		if neighbors[i].State {
			second4Alive += 1
		}
	}

	return nil
}
