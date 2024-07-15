package utils

import (
	"os"
	"image"
	"image/png"
)

func WritePNG(img image.Image, filepath string) error {
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	return png.Encode(file, img)
}
