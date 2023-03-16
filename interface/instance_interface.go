package main

import "fmt"

type Vehicle interface {
	accelerates()
	brakes()
}

func main() {

	var a Vehicle

	fmt.Println(a)

	fmt.Printf("Default Type of Interface: %T\n", a)
	fmt.Printf("Zero Value of Interface: %v\n", a)

}
