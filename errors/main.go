package main

import (
	"fmt"
	"math"
)

// ErrNegativeSqrt is returned when an attempt to compute the Sqrt of a negative number
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", float64(e))
}

// Sqrt returns the square root of x.
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	const maxDelta = 0.0000001
	z := x / 2
	prevZ := 0.0

	for math.Abs(z-prevZ) > maxDelta {
		prevZ = z
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
