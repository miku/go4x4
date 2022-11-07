package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func main() {
	// 1. Declare values for width and height (e.g. w and h).

	// 2. Declare and initialize a two-dimensional slice of uint8 to hold grayscale values.

	// 3. Fill the slice with values. You can use a fixed value (e.g. 0) or
	// calculate a value in some interesting way (e.g. "i * j % 255")

	// Move data into an image and write it out as png.
	img := image.NewRGBA(image.Rectangle{
		image.Point{0, 0},
		image.Point{w, h}})

	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			img.SetRGBA(i, j, color.RGBA{
				data[i][j],
				data[i][j],
				data[i][j],
				0xff,
			})
		}
	}
	f, err := os.Create("image.png")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	png.Encode(f, img)
}
