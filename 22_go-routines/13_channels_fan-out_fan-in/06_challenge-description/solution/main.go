package main

import (
	"fmt"
	"sync"
)

func main() {

	in := gen()

	// Fan-Out
	f0 := factorial(in)
	f1 := factorial(in)
	f2 := factorial(in)
	f3 := factorial(in)
	f4 := factorial(in)
	f5 := factorial(in)
	f6 := factorial(in)
	f7 := factorial(in)
	f8 := factorial(in)
	f9 := factorial(in)

	// Fan-In
	var count int
	for n := range merge(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9) {
		// fmt.Println(n)
		count++
		fmt.Println(count, "\t", n)
	}
}

// Fan-In
// multiplex multiple channels onto a single channel
func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// Start an output goroutine for each input channel in cs.
	// output copies values from c to out until c is closed, then calls wg.Done.
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(cs))
	// fan-in from every channel to the single output channel
	for _, c := range cs {
		go output(c)
	}
	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func gen() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10000; i++ {
			for j := 3; j < 13; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
}

func factorial(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- fact(n)
		}
		close(out)
	}()
	return out
}

func fact(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}

/*
CHALLENGE #1:
-- Change the code above to execute 1,000 factorial computations concurrently and in parallel.
-- Use the "fan out / fan in" pattern

CHALLENGE #2:
WATCH MY SOLUTION BEFORE DOING THIS CHALLENGE #2
-- While running the factorial computations, try to find how much of your resources are being used.
-- Post the percentage of your resources being used to this discussion: https://goo.gl/BxKnOL
*/
