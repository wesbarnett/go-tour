package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	z_old := 1e6
	eps := 1e-6
	for math.Abs(z_old - z) > eps {
		z_old = z
	    z -= (z*z - x) / (2*z)
	}
	return z
}

func main() {
	x := 2.0
	fmt.Println(Sqrt(x))
	fmt.Println(math.Sqrt(x))
}
