package main

import (
	"fmt"
	"math/rand"
	"sync"
)

const numFactorials = 100
const rdLimit = 20

func main() {
	var w sync.WaitGroup
	w.Add(numFactorials)
	factorial(&w)
	w.Wait()

}

func factorial(wmain *sync.WaitGroup) {
	var w sync.WaitGroup
	rand.Seed(42)

	w.Add(numFactorials + 1) // why adding this again to another WaitGroup?

	for j := 1; j <= numFactorials; j++ {

		go func() {
			f := rand.Intn(rdLimit)
			w.Wait() // wait for all being initialized
			total := 1
			for i := f; i > 0; i-- {

				total *= i
			}
			fmt.Println(f, total) // no channels used here, so we print now
			(*wmain).Done()       // one thread done
			//out <- total

		}()
		w.Done()
	}
	fmt.Println("All done with initialization, starting with calculations...")
	w.Done() // this is the +1 wait thread
}
