package randomwalk

import (
	"fmt"

	"github.com/jcocozza/gotomata/core"
)

func Main(width, height, steps int, initCoord core.Coordinate) {
	rw := RandomWalk(width, height, steps)
	rw.Grid.SetCell(true, initCoord)

	for i := 0; i < steps; i++ {
		initCoord = rw.StepHead(initCoord)
		RandomWalkToTimage(rw, initCoord, fmt.Sprintf("images/%d.png", i))
	}
}
