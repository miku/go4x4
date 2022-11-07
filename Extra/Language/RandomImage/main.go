package main

import (
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
)

func randomColor() color.Color {
	return color.RGBA{
		uint8(rand.Intn(255)),
		uint8(rand.Intn(255)),
		uint8(rand.Intn(255)),
		uint8(rand.Intn(255)),
	}
}

func main() {
	var (
		width  = 200
		height = 100

		upLeft   = image.Point{0, 0}
		lowRight = image.Point{width, height}
		img      = image.NewRGBA(image.Rectangle{upLeft, lowRight})
	)

	// Colors are defined by Red, Green, Blue, Alpha uint8 values.
	// cyan := color.RGBA{100, 200, 200, 0xff}

	// Set color for each pixel.
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, randomColor())
		}
	}

	// Encode as PNG.
	f, _ := os.Create("image.png")
	png.Encode(f, img)
}
