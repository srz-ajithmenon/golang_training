package main

import (
	"fmt"
)

func main() {

	num := []string{"Python", "Java", "Golang", "JavaScript", "Swift"}

	fmt.Println(num)
	fmt.Println("Length : ", len(num))
	fmt.Println("capacity : ", cap(num))

}
