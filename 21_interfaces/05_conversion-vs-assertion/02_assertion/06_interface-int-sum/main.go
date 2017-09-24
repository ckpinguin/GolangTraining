package main

import "fmt"

func main() {
	var val interface{} = 7
	// we need val as type int for this expression only
	fmt.Println(val.(int) + 6)
}
