package visualize

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

var gray = color.Gray{Y: 150}
var white = color.Gray{Y: 225}

func CreateImage(width, height int, filepath string, data [][]bool) *image.Gray {

	height = len(data)
	width = len(data[0])

	img := image.NewGray(image.Rect(0, 0, width, height))

	index := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if data[y][x] {
				img.SetGray(x, y, gray)
			} else {
				img.SetGray(x, y, white)
			}
			index++
		}
	}

	// Create a new PNG file to save the image
	file, err := os.Create(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Encode the grayscale image to PNG and save it to the file
	err = png.Encode(file, img)
	if err != nil {
		panic(err)
	}

	return img
}
