package main

import "fmt"

func main() {

	nums := []int{2, 5, 6, 7}

	slicePointer(&nums)
	fmt.Println(nums)

}

func slicePointer(s *[]int) {

	fmt.Println(*s)
	*s = append(*s, 9)

}
