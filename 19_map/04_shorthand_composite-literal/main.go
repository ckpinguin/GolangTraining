package main

import "fmt"

func main() {

	myGreeting := map[string]string{}
	fmt.Println(myGreeting)
	fmt.Println(myGreeting == nil)
	myGreeting["Tim"] = "Good morning."
	myGreeting["Jenny"] = "Bonjour."

	fmt.Println(myGreeting)
}
