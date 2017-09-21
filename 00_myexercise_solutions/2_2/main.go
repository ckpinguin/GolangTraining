package main

import "fmt"

func main() {
	half := func(n int) (int, bool) {
		isEven := n%2 == 0
		return n / 2, isEven
	}
	fmt.Println(half(1))
	fmt.Println(half(2))
}
