package main

import (
	"fmt"
	"time"
)

func apple() {

	for i := 1; i <= 5; i++ {
		fmt.Println(i, " apple")
	}

}

func ball() {

	for i := 1; i <= 5; i++ {
		fmt.Println(i, " ball")
	}

}

func main() {

	go apple() //goroutine
	time.Sleep(time.Microsecond * 10)
	ball()

}
