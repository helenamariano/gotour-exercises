package main

import (
	"fmt"
	"math"
)

// Sqrt returns the square root of x.
func Sqrt(x float64) float64 {
	const maxDelta = 0.0000001
	z := x / 2
	prevZ := 0.0

	for math.Abs(z-prevZ) > maxDelta {
		prevZ = z
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	x := float64(1236585)
	mine := Sqrt(x)
	lib := math.Sqrt(x)
	diff := math.Abs(mine - lib)
	fmt.Printf("Mine: %v\n", mine)
	fmt.Printf("Lib: %v\n", lib)
	fmt.Printf("Diff: %v\n", diff)
}
