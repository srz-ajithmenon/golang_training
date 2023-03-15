package main

import "fmt"

func fact1(number int) int {

	if number == 0 || number == 1 {
		return 1
	} else if number < 0 {
		fmt.Println("Invalid number")
		return number
	} else {
		return number * fact2(number-1)
	}

}

func fact2(number int) int {

	if number == 0 || number == 1 {
		return 1
	} else if number < 0 {
		fmt.Println("Invalid number")
		return number
	} else {
		return number * fact1(number-1)
	}

}

func main() {

	factorial := fact1(10)
	fmt.Println(factorial)

}
