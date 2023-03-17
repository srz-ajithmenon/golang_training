package main

import (
	"fmt"

	"sync"

	"time"
)

func bike() {

	for i := 1; i <= 5; i++ {
		time.Sleep(time.Millisecond * 10)
		fmt.Println(i, "R15")
	}

}

func car() {

	for i := 1; i <= 5; i++ {
		time.Sleep(time.Millisecond * 100)
		fmt.Println(i, "Audi R8")
	}

}

func main() {

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		bike()
		wg.Done()
	}()

	go func() {
		car()
		wg.Done()
	}()

	wg.Wait()

}
