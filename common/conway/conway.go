package conway

import (
	"os"
	"fmt"
	"image"
	"image/color"
	"image/png"

	"github.com/jcocozza/gotomata/core"
)

func ConwayGameOfLife(width, height, steps int) *core.CellularAutomata[bool] {
	grid := ConwayGrid(width, height)
	return &core.CellularAutomata[bool]{
		Grid: grid,
		RuleSet: ConwayRuleSet,
		Steps: steps,
	}
}

func PrintCGOL(cgol *core.CellularAutomata[bool]) {
	s := ""
	width := cgol.Grid.Dimensions[0]
	height := cgol.Grid.Dimensions[1]

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

func CGOLToImage(cgol *core.CellularAutomata[bool], filepath string) *image.Gray {
	var gray = color.Gray{Y: 150}
	var white = color.Gray{Y: 225}

	width := cgol.Grid.Dimensions[0]
	height := cgol.Grid.Dimensions[1]

	img := image.NewGray(image.Rect(0,0, width, height))

	idx := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			coord := []int{y,x}
			cell := cgol.Grid.GetCell(coord)
			if cell.State {
				img.SetGray(x, y, gray)
			} else {
				img.SetGray(x, y, white)
			}
			idx ++
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
