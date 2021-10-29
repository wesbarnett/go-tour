package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct{
	size int
}

func (im Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (im Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, im.size, im.size)
}

func (im Image) At(x, y int) color.Color {
	return color.RGBA{im.transform(x, y), im.transform(x, y), 255, 255}
}

func (im Image) transform(x, y int) uint8 {
	return uint8(x^y)
}

func main() {
	m := Image{256}
	pic.ShowImage(m)
}
