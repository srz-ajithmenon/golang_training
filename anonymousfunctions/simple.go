package main

import (
	"fmt"
	"math"
)

var c = func(r float64) float64 {
	return 2 * math.Pi * r
}

func main() {

	func(num1, num2, num3 int) {

		total := num3 * (num1 + num2)
		fmt.Println("Total:", total)

	}(5, 3, 2)

	fmt.Println(c(10))

}
