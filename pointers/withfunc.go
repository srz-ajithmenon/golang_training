package main

import "fmt"

func pointerAdd(a *int) {

	fmt.Println(*a + *a)

}

func main() {

	var l int = 4
	pointerAdd(&l)

}
