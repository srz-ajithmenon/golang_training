package main

import "fmt"

var name string = "vishnu"

func main() {

	var employee string
	var amount int
	var choice bool
	var measurement float32

	fmt.Println("golbally declared variable : ", name)

	fmt.Println("default value of employee (string) : ", employee)
	fmt.Println("default value of amount (int) : ", amount)
	fmt.Println("default value of choice (boolean) : ", choice)
	fmt.Println("default value of measurement (float) :", measurement)

	employee = "ajith"
	amount = 100
	choice = true
	measurement = 54.5

	fmt.Println("assigned value of employee : ", employee)
	fmt.Println("assigned value of amount : ", amount)
	fmt.Println("assigned value of choice : ", choice)
	fmt.Println("assigned value of measurement : ", measurement)

}
