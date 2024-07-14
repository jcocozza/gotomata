package randomwalk

import (
	"fmt"

	"github.com/jcocozza/gotomata/core"
)

func MainRandomWalk(width, height, steps int) {
	rw := RandomWalk(width, height, steps)
	initCoord := core.Coordinate{width/2, height/2}
	rw.Grid.SetCell(true, initCoord)

	for i := 0; i < steps; i++ {
		initCoord = rw.StepHead(initCoord)
		RandomWalkToTimage(rw, initCoord, 5, fmt.Sprintf("images/%d.png", i))
	}
}
