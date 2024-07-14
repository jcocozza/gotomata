package main

import (
	"fmt"
	"math/rand"
	"net/http"
	_ "net/http/pprof"

	"github.com/jcocozza/gotomata/common/conway"
	"github.com/jcocozza/gotomata/common/dim3"

	//"github.com/jcocozza/gotomata/common/crystal3d"
	"github.com/jcocozza/gotomata/common/crystals"
	"github.com/jcocozza/gotomata/common/elementary"
	randomwalk "github.com/jcocozza/gotomata/common/randomWalk"
	"github.com/jcocozza/gotomata/core"
)

func initSphere(width, height, depth, aliveState int, c *core.CellularAutomata[int]) {
	radius := 5// Define the radius of the sphere
	center := core.Coordinate{width / 2, height / 2, depth / 2}

	for x := center[0] - radius; x <= center[0] + radius; x++ {
		for y := center[1] - radius; y <= center[1] + radius; y++ {
			for z := center[2] - radius; z <= center[2] + radius; z++ {
				if (x-center[0])*(x-center[0])+(y-center[1])*(y-center[1])+(z-center[2])*(z-center[2]) <= radius*radius {
					p := rand.Float64()
					if p < .3 {
						c.Grid.SetCell(aliveState, core.Coordinate{x, y, z})
					}
				}
			}
		}
	}
}

func initCube(width, height, depth, aliveState int, c *core.CellularAutomata[int]) {
	center := core.Coordinate{width / 2, height / 2, depth / 2}

	sideLength := 5// Define the side length of the cube

	halfSide := sideLength / 2
	for x := center[0] - halfSide; x <= center[0]+halfSide; x++ {
		for y := center[1] - halfSide; y <= center[1]+halfSide; y++ {
			for z := center[2] - halfSide; z <= center[2]+halfSide; z++ {
				p := rand.Float64()
				if p < .55 {
					c.Grid.SetCell(aliveState, core.Coordinate{x, y, z})
				}
			}
		}
	}
}

func main() {
	go func() {
		err := http.ListenAndServe("localhost:6060", nil)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}()

	Amoebamain()
	//r678678main()
	//r445main()
	//crystal3dmain()
	//amoebamain()
	//crystalmain()
	//conwaymain()
	//randomwalkmain()
	//elementarymain()
}

func Amoebamain() {
	width := 3000
	height := 3000
	depth := 3000
	steps := 100
	c := dim3.Amoeba(width, height, depth, steps)
	initCube(width, height, depth, 4, c)
	dim3.VisualizeDim3(c)
}

func r678678main() {
	width := 3000
	height := 3000
	depth := 3000
	steps := 100
	c := dim3.R678678(width, height, depth, steps)
	initCube(width, height, depth, 3/*alive state*/, c)
	dim3.VisualizeDim3(c)
}

func r445main() {
	width := 3000
	height := 3000
	depth := 3000
	steps := 100

	center := core.Coordinate{width / 2, height / 2, depth / 2}
	c := dim3.R445(width, height, depth, steps)

	sideLength := 12 // Define the side length of the cube

	halfSide := sideLength / 2
	for x := center[0] - halfSide; x <= center[0]+halfSide; x++ {
		for y := center[1] - halfSide; y <= center[1]+halfSide; y++ {
			for z := center[2] - halfSide; z <= center[2]+halfSide; z++ {
				p := rand.Float64()
				if p < .3 {
					c.Grid.SetCell(4, core.Coordinate{x, y, z})
				}
			}
		}
	}
	dim3.VisualizeDim3(c)
}

func crystal3dmain() {
	width := 300
	height := 300
	depth := 300
	steps := 100
	c := dim3.Crystal(width, height, depth, steps)
	c.Grid.SetCell(1, core.Coordinate{width / 2, height / 2, depth / 2})
	c.Grid.SetCell(1, core.Coordinate{width/2 + 1, height / 2, depth / 2})
	c.Grid.SetCell(1, core.Coordinate{width/2 + 2, height / 2, depth / 2})
	dim3.VisualizeDim3(c)
}

func crystalmain() {
	width := 1500
	height := 1500
	steps := 200

	crystal := crystals.Crystals(width, height, steps)
	crystal.Grid.SetCell(true, core.Coordinate{0, 0})

	crystals.CrystalToImage(crystal, fmt.Sprintf("_images/%d.png", 0))
	for i := 1; i < steps; i++ {
		fmt.Printf("Step: %d/%d\n", i, steps)
		crystal.Stepp()
		crystals.CrystalToImage(crystal, fmt.Sprintf("_images/%d.png", i))
	}
}

func conwaymain() {
	width := 200
	height := 200
	steps := 1000

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
		//		conway.PrintCGOL(cgol)
		conway.CGOLToImage(cgol, fmt.Sprintf("_images/%d.png", i), 10)
	}
}

func seedsmain() {
	width := 100
	height := 100
	steps := 250

	initConfig := []core.Coordinate{
		{width / 2, height/2 - 2}, {width / 2, height / 2}, {width / 2, height/2 + 2},
	}
	seeds := conway.Seeds(width, height, steps)
	for _, coord := range initConfig {
		seeds.Grid.SetCell(true, coord)
	}

	for i := 0; i < steps; i++ {
		fmt.Printf("Step: %d/%d\n", i, steps)
		seeds.Stepp()
		//		conway.PrintCGOL(seeds)
		conway.CGOLToImage(seeds, fmt.Sprintf("_images/%d.png", i), 5)
	}
}

func randomwalkmain() {
	width := 100
	height := 100
	steps := 10000

	rw := randomwalk.RandomWalk(width, height, steps)
	initConfig := []core.Coordinate{{width / 2, height / 2}}
	for _, coord := range initConfig {
		rw.Grid.SetCell(true, coord)
	}

	initCoord := initConfig[0]
	for i := 0; i < steps; i++ {
		fmt.Printf("Step: %d/%d\n", i, steps)
		initCoord = rw.StepHead(initCoord)
		randomwalk.RandomWalkToTimage(rw, initCoord, 5, fmt.Sprintf("_images/%d.png", i))
	}
}

func elementarymain() {
	gLen := 200
	steps := 50

	ecaParr := elementary.ElementaryCellularAutomata(30, gLen, steps)
	ecaParr.Grid.SetCell(true, []int{gLen / 2})
	elementary.PrintECA(ecaParr)

	for i := 0; i < steps; i++ {
		ecaParr.Stepp()
		elementary.PrintECA(ecaParr)
	}
}
