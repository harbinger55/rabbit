package main

import "fmt"

func helloWorld() (hello string) {
	return "Hello World"
}
func main() {
	hello := helloWorld()
	fmt.Println(hello)
}