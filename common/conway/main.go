package conway

import "github.com/jcocozza/gotomata/core"

func MainConway(width, height, steps int, initConfig []core.Coordinate) {
	cgol := ConwayGameOfLife(width, height, steps)
	for _, coord := range initConfig {
		cgol.Grid.SetCell(true, coord)
	}
	for i := 0; i < steps; i++ {
		cgol.Stepp()
		PrintCGOL(cgol)
	}
}

func MainSeeds(width, height, steps int, initConfig []core.Coordinate) {
	seeds := Seeds(width, height, steps)
	for _, coord := range initConfig {
		seeds.Grid.SetCell(true, coord)
	}
	for i := 0; i < steps; i++ {
		seeds.Stepp()
		PrintCGOL(seeds)
	}
}
