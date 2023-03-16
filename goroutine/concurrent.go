package main

import (
	"fmt"
	"runtime"
	"time"
)

func count(item string) {

	for i := 1; i <= 5; i++ {
		time.Sleep(time.Millisecond * 10)
		fmt.Println(i, item)
	}

}

func main() {

	go func() { // Anonymous goroutine function
		count("apple")
	}()

	go count("ball")

	fmt.Println("No, of Goroutines : ", runtime.NumGoroutine())
	fmt.Println("No, of CPU : ", runtime.NumCPU())

	time.Sleep(time.Second)

}
