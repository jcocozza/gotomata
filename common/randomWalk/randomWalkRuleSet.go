package randomwalk

import (
	"github.com/jcocozza/gotomata/core"
	"math/rand"
)

// pick a neighbor randomly as the next place to go and set that state to true
func RandomWalkRuleSet(cell *core.Cell[bool], neighbors []*core.Cell[bool]) *core.Cell[bool] {
	index := rand.Intn(len(neighbors))
	chosenNext := neighbors[index]
	return &core.Cell[bool]{State: true, Coordinate: chosenNext.Coordinate}
}
