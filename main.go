package main

import (
	"github.com/jcocozza/gotomata/core/elementary"
	"github.com/jcocozza/gotomata/visualize"
)


func main() {

    gLen := 500
    gWidth := 500

    initLayer := make([]bool, gWidth) 
    for i := 0; i < gWidth; i++ {
        if i == gWidth - 2 {
            initLayer[i] = true 
        } else {
            initLayer[i] = false
        }
    }

    cells := elementary.NewElementaryCellularAutomata(184, initLayer, gLen, gWidth)
    cells.Run(initLayer)

    /*
    for _, row := range cells.Cells {
        printrow(row)
    }
    */

    visualize.CreateImage(400,300, cells.Cells)
}
