package main

import (
	"fmt"
)

func main() {

	convert := func(gram float64) float64 {
		return gram / 1000
	}

	fmt.Println(convert(1200))

}
