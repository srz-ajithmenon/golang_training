package main

import "fmt"

func main() {

	slice1 := []int{1, 2, 3, 4, 5, 6}

	fmt.Println(slice1)
	fmt.Println(slice1[1])
	fmt.Println(slice1[3])

	slice1[3] = 8
	fmt.Println(slice1)

}
