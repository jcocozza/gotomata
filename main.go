package main

import (
	"github.com/jcocozza/gotomata/cmd"
	"github.com/jcocozza/gotomata/core/elementary"
	"github.com/jcocozza/gotomata/visualize"
)


func main() {

    eca := elementary.NewECA(30,200)
    eca.Grid.SetValue(100, true)

    cmd.Printrow(eca.Grid.Data)
    var next *elementary.ElementaryCellularAutomata
    next = eca

    data := [][]bool{}
    data = append(data, next.Grid.Data)

    for i := 0; i < 50; i++ {
        next = next.Step()
        data = append(data, next.Grid.Data)
        cmd.Printrow(next.Grid.Data)
    }

    /*
    for _, row := range cells.Cells {
        printrow(row)
    }
    */

    visualize.CreateImage(400, 300, data)
}
