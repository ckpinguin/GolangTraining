package main

import "fmt"

// tail call optimized recursive version of fibonacci is used
func fibRec(n int, ab ...int) int {
	a, b := 0, 1     // default start values (only first call)
	if len(ab) > 0 { // optional a, b here
		a, b = ab[0], ab[1]
	}
	// fmt.Printf("n: %v\t a: %v\t b: %v\t a+b: %v\n", n, a, b, a+b)
	if n > 0 {
		// fmt.Printf("sumEvenFib(%v, %v, %v)\n", n-1, b, a+b)
		return fibRec(n-1, b, a+b)
	}
	return a
}

func fibIter(n int) int {
	a, b := 0, 1
	for n > 0 {
		a, b = b, a+b
		n--
	}
	return a
}

func fibSumEvenUpTo(n int) int {
	sum := 0
	for n > 0 {
		fib := fibRec(n)
		if isEven(fib) {
			sum += fib
		}
		n--
	}
	return sum
}

func isEven(n int) bool {
	// fmt.Println(n % 2)
	return n%2 == 0
}

func main() {
	fmt.Println(fibRec(10))
	fmt.Println(fibIter(10))
	fmt.Println(fibSumEvenUpTo(50))
}
