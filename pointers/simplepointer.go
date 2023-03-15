package main

import "fmt"

func main() {

	var x int = 100

	fmt.Println(x)

	var y *int = &x
	*y = 200

	fmt.Println(x)

	k := new(int) // in-build function to declare variable as pointer
	*k = 90
	fmt.Println(k)
	fmt.Println(*k)

	p := &y
	fmt.Println(*p)  // *p == value of y == address of x
	fmt.Println(**p) // **p = *(*p) = *y = x

}
