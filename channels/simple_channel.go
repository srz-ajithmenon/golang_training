package main

import (
	"fmt"
)

func main() {

	c := make(chan string)

	go count("Peter is a good boy", c)

	msg := <-c

	fmt.Println(msg)

}

func count(person string, ch chan string) {
	ch <- person
}
