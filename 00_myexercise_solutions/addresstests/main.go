package main

import "fmt"

func main() {
	a := 1
	var b int
	const c = 3     // untyped constant
	const d int = 4 // typed constant
	fmt.Println("&a: ", &a)
	fmt.Println("&b: ", &b)
	fmt.Println("&b dereferenced: ", *(&b))
	// fmt.Println("&c: ", &c)
	// fmt.Println("&d: ", &d)
	// fmt.Println("12: ", &int(12))
}
