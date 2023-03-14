package main

import "fmt"

func main() {

	a1 := [3]int{}
	a2 := [6]int{1, 2, 3, 4}
	a3 := [5]string{"a", "b", "c", "d", "e"}
	a4 := [7]string{"a", "b", "c", "d"}

	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)

	array1 := [10]int{2: 10, 5: 20, 7: 40} // if we only know a few element (here 3rd, 6th and 8th  out of 10 elements)

	fmt.Println(array1)

	fmt.Println("length of array1 = ", len(array1)) // len(array1) : to find length of array1

}
