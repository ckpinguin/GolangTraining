package main

import "fmt"

func main() {
	const b int = 12
	a := 43

	fmt.Println("a - ", a)
	fmt.Println("a's memory address - ", &a)
	fmt.Printf("%d \n", &a)

	fmt.Println("b - ", b)
	// fmt.Println("b's memory address - ", &b)
	// fmt.Printf("%d \n", &b)
}
