package main

import "fmt"

func main() {

	var a, b = 40, 8

	fmt.Printf("a + b = %d\n", a+b)
	fmt.Printf("a - b = %d\n", a-b)
	fmt.Printf("a * b = %d\n", a*b)
	fmt.Printf("a / b = %d\n", a/b)
	fmt.Printf("a mod b = %d\n", a%b)
	fmt.Printf("a = %d\n", a)

	a++
	fmt.Printf("a++ = %d\n", a)
	fmt.Printf("b = %d\n", b)

	b--
	fmt.Printf("b-- = %d\n", b)

	n := 12.0
	m := 5.0

	fmt.Printf("%g / %g = %g\n", n, m, n/m)

}
