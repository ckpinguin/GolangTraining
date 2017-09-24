package main

import "fmt"

func main() {
	rem := 7.24
	fmt.Printf("%T\n", rem)
	fmt.Printf("%T\n", int(rem))

	var val interface{} = 7
	fmt.Printf("%T\n", val)
	// this bonks:
	fmt.Printf("%T\n", int(val))
	// we need assertion for interface types!
	// fmt.Printf("%T\n", val.(int))
}
