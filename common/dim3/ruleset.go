package dim3

import (
	"github.com/jcocozza/gotomata/core"
)

// check if a value is in a list of ints
func isIn(val int, lst []int) bool {
	for _, elm := range lst {
		if val == elm {
			return true
		}
	}
	return false
}

// assume ints, where 0 is dead and totalStates - 1 is the alive state
func GenerateRuleSet[T int](numNeighborsToSurvive []int, numNeighborsToBorn []int, totalStates int) func(cell *core.Cell[int], neighbors []*core.Cell[int]) *core.Cell[int] {
	aliveState := totalStates - 1
	deadState := 0

	var ruleSet = func(cell *core.Cell[int], neighbors []*core.Cell[int]) *core.Cell[int] {

		totalNeighbors := 0
		for _, neighbor := range neighbors {
			if neighbor.State == aliveState {
				totalNeighbors++
			}
		}

		switch cell.State {
			case aliveState:
				if isIn(totalNeighbors, numNeighborsToSurvive) {
					return &core.Cell[int]{State: aliveState, Coordinate: cell.Coordinate}
				}
				return &core.Cell[int]{State: aliveState - 1, Coordinate: cell.Coordinate}
			case deadState:
				if isIn(totalNeighbors, numNeighborsToBorn) {
					return &core.Cell[int]{State: aliveState, Coordinate: cell.Coordinate}
				}
				return &core.Cell[int]{State: deadState, Coordinate: cell.Coordinate}
			default:
				return &core.Cell[int]{State: cell.State - 1, Coordinate: cell.Coordinate}
		}
	}
	return ruleSet
}
