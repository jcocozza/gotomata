package visualize

import (
    "os"
    "image/png"
	"image"
	"image/color"
)

var gray = color.Gray{Y: 150}
var white = color.Gray{Y: 225}

func CreateImage(width, height int, data [][]bool) *image.Gray {
	
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
	file, err := os.Create("output.png")
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
