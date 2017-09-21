package main

import (
	"fmt"
)

func main() {
	m := max(20, 12, 1, 4, 9, 2, 19, 17, 11)
	fmt.Println(m)
}

func max(n ...int) int {
	max := 0
	for _, i := range n {
		if i > max {
			max = i
		}
	}
	return max
}
