package main

import (
	"fmt"
)

type Students struct {
	name string
	id   int
	Parents
}
type Parents struct {
	firstname string
}

func main() {

	st1 := Students{
		name: "Ajith",
		id:   006,
		Parents: Parents{
			firstname: "Jayachandran",
		},
	}

	fmt.Println("Student id : ", st1.name)
	fmt.Println("Student name : ", st1.id)
	fmt.Println("Parent name : ", st1.Parents.firstname) // st1.firstname  can also be used as it is unique field in the program

}
