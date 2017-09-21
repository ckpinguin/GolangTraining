package main

import "fmt"

type customer struct {
	name string
	age  int
}

func main() {
	c1 := customer{"Todd", 44}
	fmt.Println(&c1)
	fmt.Println(&c1.name) // 0x8201e4120

	changeMe(&c1)

	fmt.Println(c1)       // {Rocky 44}
	fmt.Println(&c1.name) // 0x8201e4120

	changeMe2(c1)

	fmt.Println(c1)
	fmt.Println(&c1.name)
}

func changeMe(z *customer) {
	fmt.Println(z)       // &{Todd 44}
	fmt.Println(&z.name) // 0x8201e4120
	z.name = "Rocky"
	fmt.Println(z)       // &{Rocky 44}
	fmt.Println(&z.name) // 0x8201e4120

}

func changeMe2(z customer) {
	fmt.Println(z)
	fmt.Println(&z.name)
	z.name = "Hubaaa"
	fmt.Println(z)
}
