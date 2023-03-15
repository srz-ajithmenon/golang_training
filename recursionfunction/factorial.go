package main

import "fmt"

func fact(number int) int {

	if number == 0 || number == 1 {
		return 1
	} else if number < 0 {
		fmt.Println("Invalid number")
		return number
	} else {
		return number * fact(number-1)
	}

}

func main() {

	factorial := fact(10)
	fmt.Println(factorial)

}
