package main

import (
	"fmt"
	"sync/atomic"
)

var count int64

func main() {
	c1 := incrementor("1")
	c2 := incrementor("2")
	<-c1
	<-c2
	fmt.Println("Final Counter:", count)
}

func incrementor(s string) <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 20; i++ {
			atomic.AddInt64(&count, 1)
			fmt.Println("Process: "+s+" printing:", i)
		}
		close(out)
	}()
	return out
}

/*
CHALLENGE #1
-- Take the code from above and change it to use channels instead of wait groups and atomicity
*/
