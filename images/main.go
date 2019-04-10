package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

// Image is our custom implementation of the Image interface
type Image struct{}

// ColorModel returns RGBAModel
func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

// Bounds returns the image domain
func (i Image) Bounds() image.Rectangle {
	const (
		h = 256
		w = 256
	)
	return image.Rect(0, 0, w, h)
}

// At returns the colour at pixel (x,y)
func (i Image) At(x, y int) color.Color {
	//v := uint8((x + y) / 2)
	//v := uint8(x * y)
	v := uint8(x ^ y)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
