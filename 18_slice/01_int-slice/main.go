package main

import "fmt"

func main() {
	myArr := [3]int{1, 2, 4}
	// myByteSlice := []int(myArr)
	fmt.Println(myArr)

	mySlice := []int{1, 3, 5, 7, 9, 11}
	fmt.Printf("%T\n", mySlice)
	fmt.Println(mySlice)
}
