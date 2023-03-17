package main

import (
	"fmt"

	"time"
)

func count(fruit string) {

	var i int

	for i = 1; i <= 5; i++ {
		time.Sleep(time.Millisecond * 10)
		fmt.Println(i, fruit)

	}

}

func main() {

	defer count("apple")
	count("grapes")
	count("oranges")

}
