package main

import (
	"fmt"
)

func count(n int) {

	if n > 0 {
		fmt.Println(n)
		count(n - 1)
	}

}

func main() {

	count(10)

}
