package main

import "fmt"

func main() {

	c := make(chan string, 3)

	c <- "ajith"
	c <- "Ashique"
	c <- "vishnu"
	// c <- "amal"   //  Entering more than 3 in this will make factal error

	message := <-c
	fmt.Println(message)

	c <- "aldes" //  as one was sent out we got space to add new value to channel

	message2 := <-c
	message3 := <-c

	fmt.Println(message2)
	fmt.Println(message3)
	fmt.Println(<-c)

}
