package main

import "fmt"

func total(a int, nums ...int) {

	sum := 0
	for _, e := range nums {
		sum = sum + a*e
	}

	fmt.Println(sum)

}

func main() {

	total(10, 2, 8, 6, 4)

}
