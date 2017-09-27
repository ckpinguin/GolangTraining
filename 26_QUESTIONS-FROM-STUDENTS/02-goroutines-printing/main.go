package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("In goroutine:", i)
			c <- i // blocking here, until main() takes off one
		}
		close(c)
	}()

	for n := range c {
		fmt.Println("In main: ", n)
	}
}
