package main

import (
	"fmt"
)

type Student struct {
	name  string
	id    int
	class int
}
type Parent struct {
	firstName string
	lastName  string
	childID   int
}

func (s Student) printDetails() {
	fmt.Printf("%s with ID No. %d is in class %d \n", s.name, s.id, s.class)
}
func (p Parent) printDetails() {
	fmt.Printf("%s %s is parent of child with id %d \n", p.firstName, p.lastName, p.childID)
}

func main() {

	st1 := Student{
		name:  "Ajith",
		id:    006,
		class: 2,
	}

	st1.printDetails()

	pt1 := Parent{
		firstName: "Jayachandran",
		lastName:  "M",
		childID:   006,
	}

	pt1.printDetails()

}
