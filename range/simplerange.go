package main

import "fmt"

func main() {

	nums := [...]int{15, 20, 33, 47, 55, 26}
	fmt.Println(nums)

	for i, e := range nums {

		fmt.Printf("Element %v is at index %d \n", e, i)

	}

	// we could also print same using just index
	fmt.Println("____________________________________")

	for i := range nums {

		fmt.Printf("Element %v is at index %d \n", nums[i], i)

	}

	// we could use range even on string variable
	fmt.Println("____________________________________")

	str := "你好世界人民"
	for _, element := range str {
		fmt.Printf("%c\n", element)
	}

}
