package main

import "fmt"

func main() {

	array1 := [8]int{1, 2, 3, 4, 5, 6, 7, 8}

	slice1 := array1[1:5]

	fmt.Printf("Array = %v\n", array1)

	fmt.Printf("Slice = %v\n", slice1)
	fmt.Printf("length = %d\n", len(slice1))
	fmt.Printf("capacity = %d\n", cap(slice1))

}
