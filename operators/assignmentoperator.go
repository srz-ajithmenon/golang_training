package main

import "fmt"

func main() {

	x := 30
	y := 10
	fmt.Println("x=", x)
	fmt.Println("y=", y)

	x += y
	fmt.Println("x +=y", x)

	x = 70
	fmt.Println("x=", x)

	x -= y
	fmt.Println("x=", x)

	x = 14
	fmt.Println("x=", x)

	x *= y
	fmt.Println("x=", x)

	x = 90
	fmt.Println("x=", x)

	x /= y
	fmt.Println("x=", x)

	x = 55
	fmt.Println("x=", x)

	x %= y
	fmt.Println("x=", x)

}
