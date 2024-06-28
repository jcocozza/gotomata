package main

import (
	"fmt"

//	"github.com/jcocozza/gotomata/cmd"
	"github.com/jcocozza/gotomata/core/conway"
	"github.com/jcocozza/gotomata/core/elementary"
	"github.com/jcocozza/gotomata/visualize"
)
/*
func elementaryMain() {
    rule := 30
    eca := elementary.NewECA(uint8(rule),200)
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

    visualize.CreateImage(400, 300,fmt.Sprintf("rule%d.png",rule), data)
}
*/
func elementaryMain(rule uint8) {
    width := 1000 
    eca := elementary.NewECA(rule, width)

    // initialize the center of the first row to true
    eca.Grid.SetValue(500, true)

    data := eca.Run(600)
//    cmd.PrintRows(data)
    visualize.CreateImage(800, 800,fmt.Sprintf("images/elementary/rule%d.png",rule), data)
}

func conwayMain() {
    gol := conway.NewGOL(200,200)
    gol.Grid.SetValueByCoord(true, []int{100,100})
    gol.Grid.SetValueByCoord(true, []int{101,101})
    gol.Grid.SetValueByCoord(true, []int{102,99})
    gol.Grid.SetValueByCoord(true, []int{102,100})
    gol.Grid.SetValueByCoord(true, []int{102,101})

    var next *conway.GameOfLife
    next = gol

    for i := 0; i < 591; i++ {
        next = next.Step()
        rows := next.ToRows()
        //cmd.PrintRows(rows)
        visualize.CreateImage(1000, 1000, fmt.Sprintf("images/%d.png", i), rows)
    }
}

func main() {
//    conwayMain()
//    elementaryMain(30)

/*
    for i := 1; i <= 255; i++ {
        print(fmt.Sprintf("running rule %d\n", i))
        elementaryMain(uint8(i)) 
    }
*/
}
