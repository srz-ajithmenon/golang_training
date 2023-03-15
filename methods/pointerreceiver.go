package main

import "fmt"

type Student1 struct {
	name  string
	id    int
	class int
}

func (s *Student1) changeDetails(id, class int) {
	s.id = id
	s.class = class
}

func main() {

	st1 := Student1{
		name:  "Ajith",
		id:    6,
		class: 2,
	}

	fmt.Printf("Before id = %d and class %d \n", st1.id, st1.class)

	(&st1).changeDetails(5, 5)
	//	st1.changeDetails(5, 5)

	fmt.Printf("Before id = %d and class %d \n", st1.id, st1.class)

}
