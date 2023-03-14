package main

import "fmt"

func main() {

	x := 15 //0000 1111
	y := 72 //0100 1000

	fmt.Println("x & y  =", x&y)

	fmt.Println("x | y  =", x|y)

	fmt.Println("x ^ y  =", x^y)

	fmt.Println("x << 1 =", x<<1)

	fmt.Println("x >> 1 =", x>>1)

}
