package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"

	"github.com/jcocozza/gotomata/common/conway"
	"github.com/jcocozza/gotomata/common/crystals"
	"github.com/jcocozza/gotomata/common/elementary"
	randomwalk "github.com/jcocozza/gotomata/common/randomWalk"
	"github.com/jcocozza/gotomata/core"
)

func main() {
	go func() {
	err := http.ListenAndServe("localhost:6060", nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	}()

//	conwaymain()
//	randomwalkmain()
	//elementarymain()
	seedsmain()
//	crystalmain()
}

func crystalmain() {
	width := 1000
	height := 1000
	steps :=  31

	crystal := crystals.Crystals(width, height, steps)
	crystal.Grid.SetCell(true, core.Coordinate{0, 0})

	crystals.CrystalToImage(crystal, fmt.Sprintf("images/%d.png", 0))
	for i := 1; i < steps; i++ {
		fmt.Printf("Step: %d/%d\n", i, steps)
		crystal.Stepp()
		crystals.CrystalToImage(crystal, fmt.Sprintf("images/%d.png", i))
	}
}

func conwaymain() {
	width := 100
	height := 100
	steps := 60000

	//initConfig := conway.AcornConfig(width, height)
	initConfig := conway.AcornConfig(width, height)
	cgol := conway.ConwayGameOfLife(width, height, steps)
	for _, coord := range initConfig {
		cgol.Grid.SetCell(true, coord)
	}
	//conway.PrintCGOL(cgol)
	for i := 0; i < steps; i++ {
		fmt.Printf("Step: %d/%d\n", i, steps)
		cgol.Stepp()
		conway.PrintCGOL(cgol)
//		conway.CGOLToImage(cgol, fmt.Sprintf("images/%d.png", i))
	}
}

func seedsmain() {
	width := 100
	height := 100
	steps := 250

	initConfig := []core.Coordinate{
		{width/2, height/2 - 2}, {width/2, height/2},{width/2, height/2 + 2},
	}
	seeds := conway.Seeds(width, height, steps)
	for _, coord := range initConfig {
		seeds.Grid.SetCell(true, coord)
	}

	for i := 0; i < steps; i++ {
		fmt.Printf("Step: %d/%d\n", i, steps)
		seeds.Stepp()
//		conway.PrintCGOL(seeds)
		conway.CGOLToImage(seeds, fmt.Sprintf("images/%d.png", i))
	}
}

func randomwalkmain() {
	width := 100
	height := 100
	steps := 10000

	rw := randomwalk.RandomWalk(width, height, steps)
	initConfig := []core.Coordinate{{width/2, height/2}}
	for _, coord := range initConfig {
		rw.Grid.SetCell(true, coord)
	}

	initCoord := initConfig[0]
	for i := 0; i < steps; i++ {
		fmt.Printf("Step: %d/%d\n", i, steps)
		initCoord = rw.StepHead(initCoord)
		randomwalk.RandomWalkToTimage(rw, initCoord, fmt.Sprintf("images/%d.png", i))
	}
}

func elementarymain() {
	gLen := 100
	steps := 100

	ecaParr := elementary.ElementaryCellularAutomata(30, gLen, steps)
	ecaParr.Grid.SetCell(true, []int{gLen / 2})
	elementary.PrintECA(ecaParr)

	for i := 0; i < steps; i++ {
		ecaParr.Stepp()
		elementary.PrintECA(ecaParr)
	}
}
