package main

import (
	"fmt"
	"time"
)

func main() {

	n := 6
	go calcuate(n)
	time.Sleep(time.Second)

}

func calcuate(num int) {

	square := num * num
	fmt.Printf("Square of %d : %d ", num, square)

}
