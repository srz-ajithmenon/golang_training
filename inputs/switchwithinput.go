package main

import "fmt"

func main() {

	var name string

	fmt.Println("Enter your first name: ")
	fmt.Scan(&name)

	switch name {

	case "Ajith":
		fmt.Println("Hi Ajith")

	case "Vishnu":
		fmt.Println("Hi Vishnu")

	default:
		fmt.Println("No Match")

	}

}
