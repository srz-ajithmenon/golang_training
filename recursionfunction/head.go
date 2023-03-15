package main

import (
	"fmt"
)

func count(n int) {

	if n > 0 {
		count(n - 1)
		fmt.Println(n)
	}

}

func main() {

	count(10)

}
