package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func main() {
	w, h := 480, 480
	var data [][]uint8 = make([][]uint8, w)

	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			// data[i] = append(data[i], uint8(i*j%255))
			data[i] = append(data[i], uint8(8*(i+j)*i+30*j))
		}
	}
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
