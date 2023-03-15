package main

import "fmt"

func main() {

	var a int = 5052
	var b int = 30
	var c int = 9000

	maximum := maxi(a, b, c)
	fmt.Printf("The maximum value is: %d\n", maximum)

}

func maxi(num1, num2, num3 int) int {

	var result int

	if num1 > num2 && num1 > num3 {
		result = num1
	} else if num2 > num3 {
		result = num2
	} else {
		result = num3
	}

	return result

}
