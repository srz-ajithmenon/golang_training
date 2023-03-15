package main

import "fmt"

type Student2 struct {
	name  string
	id    int
	class int
}

func (s Student2) changeDetails(id, class int) {
	s.id = id
	s.class = class
}

func main() {

	st1 := Student2{
		name:  "Ajith",
		id:    006,
		class: 2,
	}

	fmt.Printf("Before id = %d and class %d \n", st1.id, st1.class)

	st1.changeDetails(005, 5) // changes will not refelct in main

	fmt.Printf("Before id = %d and class %d \n", st1.id, st1.class)

}
