package main

import "github.com/jcocozza/gotomata/common/elementary"


func main() {
    gLen := 200
    steps := 50
    //gWidth := 500

    eca := elementary.ElementaryCellularAutomata(30, gLen, steps)
    eca.Grid.SetCell(true, []int{gLen/2})
    elementary.PrintECA(eca)
    for i := 0; i < steps; i ++ {
        eca.Step()
        elementary.PrintECA(eca)
    }
}
