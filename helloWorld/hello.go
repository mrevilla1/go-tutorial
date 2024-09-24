package main

import (
	"fmt"

	"rsc.io/quote"
)

func HelloWorld() {
	fmt.Println("Hello World!!")
	fmt.Println(quote.Glass())
}
func main() {
	HelloWorld()
}
