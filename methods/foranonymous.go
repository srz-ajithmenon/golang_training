package main

import (
	"fmt"
)

type Students struct {
	name  string
	id    int
	class int
	Parents
}
type Parents struct {
	firstName string
	lastName  string
}

func (p Parents) parentDetails() {
	fmt.Printf("name of parent is %s %s \n", p.firstName, p.lastName)
}

func main() {

	st1 := Students{
		name:  "Ajith",
		id:    006,
		class: 2,
		Parents: Parents{
			firstName: "Jayachandran",
			lastName:  "M",
		},
	}

	st1.parentDetails() // st1.Parents.parentDetails()

	//  even though parentDetails() is created for 'Parents', it canbe used with student as Parents is Anonymous field of Students

}
