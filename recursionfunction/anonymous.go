package main

import (
	"fmt"
)

func main() {

	var x func(int)

	x = func(n int) {

		if n > 0 {
			fmt.Println(n)
			x(n - 1)
		}

	}

	x(10)

}
