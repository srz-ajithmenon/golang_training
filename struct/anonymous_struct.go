package main

import (
	"fmt"
)

func main() {

	u1 := struct {
		name   string
		age    int
		salary float32
	}{"ajith", 24, 10000}

	fmt.Println(u1)

}
