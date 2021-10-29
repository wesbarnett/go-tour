package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	result := make([][]uint8, dy)
	for iy := 0; iy < dy; iy += 1 {
		x := make([]uint8, dx)
		for ix := 0; ix < dx; ix += 1 {
			x[ix] = uint8(ix^iy)
		}
		result[iy] = x
	}
	return result
}

func main() {
	pic.Show(Pic)
}
