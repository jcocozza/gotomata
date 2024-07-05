package main

import (
	"fmt"
	"time"
	"net/http"
	_ "net/http/pprof"

	"github.com/jcocozza/gotomata/common/conway"
	"github.com/jcocozza/gotomata/common/elementary"
)

func main() {
	go func() {
	err := http.ListenAndServe("localhost:6060", nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	}()

	conwaymain()
}

func conwaymain() {
	width := 1000
	height := 1000
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
		//conway.PrintCGOL(cgol)
//		conway.CGOLToImage(cgol, fmt.Sprintf("images/%d.png", i))
	}
}

func elementarymain() {
	//    gLen := 500
	// one billion cells
	gLen := 1000000000
	// one million steps
	steps := 1000000
	//steps := 400
	//gWidth := 500

	ecaParr := elementary.ElementaryCellularAutomata(30, gLen, steps)
	ecaParr.Grid.SetCell(true, []int{gLen / 2})
	//    elementary.PrintECA(eca)

	fmt.Println("testing parallel compute speed...")
	startParr := time.Now()
	for i := 0; i < steps; i++ {
		stepStart := time.Now()
		ecaParr.Stepp()
		stepEnd := time.Now()
		totalStep := stepEnd.Sub(stepStart)
		fmt.Printf("Parallel Step %d/%d completed in: %f\n", i, steps, totalStep.Minutes())
		//        elementary.PrintECA(eca)
	}
	endParr := time.Now()
	totalParr := endParr.Sub(startParr)

	fmt.Println("Parallel total time:", totalParr.Seconds())

	ecaLinear := elementary.ElementaryCellularAutomata(30, gLen, steps)
	ecaLinear.Grid.SetCell(true, []int{gLen / 2})
	fmt.Println("testing linear compute speed...")
	startLinear := time.Now()
	for i := 0; i < steps; i++ {
		stepStart := time.Now()
		ecaParr.Step()
		stepEnd := time.Now()
		totalStep := stepEnd.Sub(stepStart)
		fmt.Printf("Linear Step %d/%d completed in: %f\n", i, steps, totalStep.Minutes())
		//        elementary.PrintECA(eca)
	}
	endLinear := time.Now()
	totalLinear := endLinear.Sub(startLinear)
	fmt.Println("Linear total time:", totalLinear.Seconds())
}
