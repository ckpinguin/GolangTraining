package main

import "fmt"

func main() {
	fmt.Println(half(1))
	fmt.Println(half(2))
}

func half(n int) (int, bool) {
	isEven := n%2 == 0
	return n / 2, isEven
}
