package crystals

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"

	"github.com/jcocozza/gotomata/core"
)

const (
	hexSize = 5
)

func Crystals(width, height, steps int) *core.CellularAutomata[bool] {
	grid := CrystalGrid(width, height)
	return &core.CellularAutomata[bool]{
		Grid:    grid,
		RuleSet: CrystalRuleSet,
		Steps:   steps,
	}
}

func CrystalToImage(crystal *core.CellularAutomata[bool], filepath string) *image.RGBA {
	width := crystal.Grid.Dimensions[0]
	height := crystal.Grid.Dimensions[1]

	totDraw := 0
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for _, key := range crystal.Grid.Cells.GetAllKeys() {
		cell := crystal.Grid.GetCellByHash(key)
		if cell.State {
			drawHex(img, cell.Coordinate[0], cell.Coordinate[1], width, height, color.RGBA{255, 0, 0, 255})
			totDraw++
		}
	}

	// Save the image to a file
	file, err := os.Create(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	png.Encode(file, img)

	return img
}

func drawHex(img *image.RGBA, q, r, width, height int, c color.Color) {
	centerX, centerY := hexToPixel(q, r, width, height)

	for x := centerX - hexSize; x <= centerX+hexSize; x++ {
		for y := centerY - hexSize; y <= centerY+hexSize; y++ {
			if isInsideHex(float64(x-centerX), float64(y-centerY)) {
				img.Set(x, y, c)
			}
		}
	}
}

func hexToPixel(q, r, width, height int) (x, y int) {
	x = width/2 + int(float64(hexSize)*(math.Sqrt(3)*float64(q)+math.Sqrt(3)/2*float64(r)))
	y = height/2 + int(float64(hexSize)*(3./2*float64(r)))
	return
}

func isInsideHex(x, y float64) bool {
	return math.Abs(x) < float64(hexSize)*math.Sqrt(3)/2 &&
		math.Abs(y) < float64(hexSize) &&
		math.Abs(x)*math.Sqrt(3)/3+math.Abs(y) < float64(hexSize)
}
