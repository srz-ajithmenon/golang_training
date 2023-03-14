package main

import (
	"fmt"
)

func main() {

	var x, y, z int = 20, 30, 50

	fmt.Println("x", x)
	fmt.Println("y", y)
	fmt.Println("z", z)

	var (
		a int = 20
		b     = 30
		c     = 50
		d int
	)

	fmt.Println("a", a)
	fmt.Println("b", b)
	fmt.Println("c", c)
	fmt.Println("d", d)

}
