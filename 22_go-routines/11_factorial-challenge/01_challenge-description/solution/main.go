package main

import (
	"fmt"
)

func main() {
	c := factorial(65)
	for n := range c {
		fmt.Println(n)
	}
}

func factorial(n uint64) chan uint64 {
	out := make(chan uint64)

	go func() {
		acc := uint64(1)
		for i := n; i > 0; i-- {
			acc *= i
		}
		out <- acc
		close(out)
	}()
	return out
}

/*
CHALLENGE #1:
-- Use goroutines and channels to calculate factorial

CHALLENGE #2:
-- Why might you want to use goroutines and channels to calculate factorial?
---- Formulate your own answer, then post that answer to this discussion area: https://goo.gl/flGsyX
---- Read a few of the other answers at the discussion area to see the reasons of others
*/
