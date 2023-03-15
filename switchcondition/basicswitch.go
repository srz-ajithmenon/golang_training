package main

import "fmt"

func main() {

	name := "amal"

	switch name {

	case "Ajith":
		fmt.Println("Hi Ajith")

	case "Vishnu":
		fmt.Println("Hi vishnu")

	case "Deepthi":
		fmt.Println("Hi Deepthi")

	case "Ashique":
		fmt.Println("Hi Ashique")
		fallthrough

	case "Vidhu":
		fmt.Println("Hi Vidhu")

	default:
		fmt.Println("Name don't match")
	}

}
