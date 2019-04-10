package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	fibNMinus1 := 1
	fibNMinus2 := 0
	n := 0

	return func() int {
		defer func() {
			n++
		}()

		if n < 2 {
			return n
		}

		fibN := fibNMinus1 + fibNMinus2
		fibNMinus2 = fibNMinus1
		fibNMinus1 = fibN
		return fibN
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
