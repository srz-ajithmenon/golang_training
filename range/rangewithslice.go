package main

import "fmt"

func main() {

	nums := make([]int, 5)
	fmt.Println("Before : ", nums)

	for i := range nums {
		nums[i] = i
	}
	fmt.Println("After : ", nums)

}
