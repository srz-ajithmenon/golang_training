package main

import "fmt"

func main() {

	n1 := make([]int, 2, 6)

	fmt.Printf("n1 = %v\n", n1)
	fmt.Printf("length_n1 = %d\n", len(n1))
	fmt.Printf("capacity_n1 = %d\n", cap(n1))

	n2 := make([]int, 4)

	fmt.Printf("n2 = %v\n", n2)
	fmt.Printf("length_n2 = %d\n", len(n2))
	fmt.Printf("capacity_n2 = %d\n", cap(n2))

}
