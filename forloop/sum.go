package main

import "fmt"

func main() {

	n, sum := 5, 0

	for i := 1; i <= n; i++ {

		sum += i

	}

	fmt.Println("sum = ", sum)

}
