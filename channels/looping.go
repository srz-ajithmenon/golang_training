package main

import (
	"fmt"
)

func main() {

	c := make(chan string)
	go count("Peter is a good boy", c)

	for msg := range c { // loop through channel until it become empty or closed
		fmt.Println(msg)
	}

}

func count(person string, ch chan string) {

	for i := 1; i <= 6; i++ {
		ch <- person
	}
	close(ch) // used to close channel ( sending channel ) // fatal error resolved
}
