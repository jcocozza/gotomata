package main

import (
	"fmt"
	"time"

	"github.com/jcocozza/gotomata/common/elementary"
)


func main() {
//    gLen := 500
    // one billion cells
    gLen := 1000000000
    // one million steps
    steps := 1000000
   //steps := 400
    //gWidth := 500

    ecaParr := elementary.ElementaryCellularAutomata(30, gLen, steps)
    ecaParr.Grid.SetCell(true, []int{gLen/2})
//    elementary.PrintECA(eca)

    fmt.Println("testing parallel compute speed...")
    startParr := time.Now()
    for i := 0; i < steps; i ++ {
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
    ecaLinear.Grid.SetCell(true, []int{gLen/2})
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
