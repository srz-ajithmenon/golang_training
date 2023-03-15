package main

import (
	"fmt"
)

func main() {

	t := func() func(num1, num2, num3 int) int {

		w := func(num1, num2, num3 int) int {
			total := num3 * (num1 + num2)
			return total
		}

		return w

	}

	z := t()

	fmt.Println("the total is", z(2, 3, 4))

}
