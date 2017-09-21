package main

import "fmt"

// tail call optimized recursive version of fibonacci is used
func sumEvenFib(n int, ab ...int) int {
	a, b := 0, 1     // default start values (only first call)
	if len(ab) > 0 { // optional a, b here
		a, b = ab[0], ab[1]
	}
	fmt.Printf("n: %v\t a: %v\t b: %v\t a+b: %v\n", n, a, b, a+b)
	if n > 0 {
		if isEven(a + b) {
			return sumEvenFib(n-1, b, a)
		}
		return sumEvenFib(n-1, b, a+b)
	}
	return a
}

func isEven(n int) bool {
	// fmt.Println(n % 2)
	return n%2 == 0
}

func main() {
	fmt.Println(sumEvenFib(10))
}
