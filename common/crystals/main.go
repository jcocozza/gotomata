package crystals

import (
	"fmt"

	"github.com/jcocozza/gotomata/core"
)

func Main(width, height, steps int, initConfig []core.Coordinate) {
	crystal := Crystals(width, height, steps)
	for _, coord := range initConfig {
		crystal.Grid.SetCell(true, coord)
	}
	for i := 0; i < steps; i++ {
		crystal.Stepp()
		CrystalToImage(crystal, fmt.Sprintf("images/%d.png", i))
	}
}
