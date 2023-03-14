package main

import (
	"fmt"
)

func main() {

	var a = 30
	var b = "Hello"
	var c = true

	fmt.Printf("%v is an %T\n", a, a)
	fmt.Printf("%v is a %T\n", b, b)
	fmt.Printf("%v is a %T\n", c, c)

}
