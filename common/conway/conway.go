package conway

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/jcocozza/gotomata/core"
)

func ConwayGameOfLife(width, height, steps int) *core.CellularAutomata[bool] {
	grid := ConwayGrid(width, height)
	return &core.CellularAutomata[bool]{
		Grid:    grid,
		RuleSet: ConwayRuleSet,
		Steps:   steps,
	}
}

func Seeds(width, height, steps int) *core.CellularAutomata[bool] {
	grid := ConwayGrid(width, height)
	return &core.CellularAutomata[bool]{
		Grid:    grid,
		RuleSet: SeedsRuleSet,
		Steps:   steps,
	}
}

func PrintCGOL(cgol *core.CellularAutomata[bool]) {
	s := ""
	width := cgol.Grid.Dimensions[0]
	height := cgol.Grid.Dimensions[1]

	dims := cgol.Grid.ComputeMaxDims()
	testWidth, testHeight := dims[0], dims[1]
	fmt.Printf("Width: %d, Height: %d\n", width, height)
	fmt.Printf("Test Width: %d, Test Height: %d\n", testWidth, testHeight)

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			coord := []int{i, j}
			//		fmt.Println("getting cell at coord:", coord)
			cell := cgol.Grid.GetCell(coord)
			//			fmt.Println("cell state: ", cell.State)
			if cell.State {
				s += "█"
			} else {
				s += "░"
			}
		}
		fmt.Println(s)
		s = ""
	}
}

func CGOLToImage(cgol *core.CellularAutomata[bool], filepath string, scale int) *image.Gray {
	var gray = color.Gray{Y: 150}
	var white = color.Gray{Y: 225}

	width := cgol.Grid.Dimensions[0] * scale
	height := cgol.Grid.Dimensions[1] * scale

	img := image.NewGray(image.Rect(0, 0, width, height))

	for y := 0; y < cgol.Grid.Dimensions[1]; y++ {
		for x := 0; x < cgol.Grid.Dimensions[0]; x++ {
			coord := []int{y, x}
			cell := cgol.Grid.GetCell(coord)
			color := white
			if cell.State {
				color = gray
			}
			for dy := 0; dy < scale; dy++ {
				for dx := 0; dx < scale; dx++ {
					img.SetGray(x*scale+dx, y*scale+dy, color)
				}
			}
		}
	}

	file, err := os.Create(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		panic(err)
	}
	return img
}
