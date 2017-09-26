package main

import (
	"fmt"
	"sync"
)

func main() {

	in := gen()

	// FAN OUT
	// Multiple functions reading from the same channel until that channel is closed
	// Distribute work across multiple functions (ten goroutines) that all read from in.
	xc := fanOut(in, 10)

	// FAN IN
	// multiplex multiple channels onto a single channel
	// merge the channels from c0 through c9 onto a single channel
	for n := range fanIn(xc...) {
		fmt.Println(n)
	}

}

func gen() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			for j := 3; j < 13; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
}

// dead-lock must be caused here:
func fanOut(in <-chan int, n int) []<-chan int {
	// This seemed to cause the deadlock first
	// xc := make([]<-chan int, n)
	var xc []<-chan int // must be initialized to 0

	for i := 0; i < n; i++ {
		// This works, but is not flexible (it has to exactly match the
		// n allocated slice size)
		// xc[i] = factorial(in)
		// This seemed to cause the deadlock first
		xc = append(xc, factorial(in))
	}
	return xc
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

func fanIn(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(cs))
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

/*
CHALLENGE #1:
-- This code throws an error: fatal error: all goroutines are asleep - deadlock!
-- fix this code!
*/
