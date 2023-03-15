package main

import (
	"fmt"
)

func countDown(n int) {
	if n > 0 {

		fmt.Println(n)

		countDown(n - 1)

	}
}

func main() {

	countDown(10)

}
