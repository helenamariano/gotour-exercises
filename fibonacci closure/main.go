package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	FnMinus1 := 1
	FnMinus2 := 0
	n := 0

	return func() int {
		defer func() {
			n++
		}()

		if n < 2 {
			return n
		}

		Fn := FnMinus1 + FnMinus2
		FnMinus2 = FnMinus1
		FnMinus1 = Fn
		return Fn
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
