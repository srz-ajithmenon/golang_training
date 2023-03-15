package main

import "fmt"

func main() {

	for i := 1; i <= 5; i++ {

		fmt.Println("I am outer loop ", i)

		for j := 1; j <= 2; j++ {

			fmt.Println("I am Inner Loop")

		}

	}

}
