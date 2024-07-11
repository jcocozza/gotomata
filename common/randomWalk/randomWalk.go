package randomwalk

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/jcocozza/gotomata/core"
)

func RandomWalk(width, height, steps int) *core.CellularAutomata[bool] {
	grid := randomWalkGrid(width, height)
	return &core.CellularAutomata[bool]{
		Grid:    grid,
		RuleSet: RandomWalkRuleSet,
		Steps:   steps,
	}
}

func PrintRandomWalk(cgol *core.CellularAutomata[bool]) {
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

func RandomWalkToTimage(cgol *core.CellularAutomata[bool], current core.Coordinate, scale int, filepath string) *image.Gray {
	var gray = color.Gray{Y: 150}
	var white = color.Gray{Y: 225}

	width := cgol.Grid.Dimensions[0] * scale
	height := cgol.Grid.Dimensions[1] * scale

	img := image.NewGray(image.Rect(0, 0, width, height))

	idx := 0
	for y := 0; y < cgol.Grid.Dimensions[1]; y++ {
		for x := 0; x < cgol.Grid.Dimensions[0]; x++ {
			coord := []int{y, x}
			cell := cgol.Grid.GetCell(coord)
			if coord[0] == current[0] && coord[1] == current[1] {
				for dy := 0; dy < scale; dy++ {
					for dx := 0; dx < scale; dx++ {
						img.Set(x*scale+dx, y*scale+dy, color.RGBA{R: 255, G: 0, B: 0, A: 255})
					}
				}
			} else {
				if cell.State {
					for dy := 0; dy < scale; dy++ {
						for dx := 0; dx < scale; dx++ {
							img.Set(x*scale+dx,y*scale+dy, gray)
						}
					}
				} else {
					for dy := 0; dy < scale; dy++ {
						for dx := 0; dx < scale; dx++ {
							img.Set(x*scale+dx,y*scale+dy, white)
						}
					}
				}
			}
			idx++
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
