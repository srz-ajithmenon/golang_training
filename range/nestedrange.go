package main

import "fmt"

func main() {

	names := []string{"ajith", "vishnu"}
	fruits := []string{"apple", "orange", "banana"}

	for _, name := range names {
		fmt.Printf("%s:\n", name)

		for _, fruit := range fruits {
			fmt.Println(" ", fruit)
		}

	}

}
