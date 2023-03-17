package main

import "fmt"

type Employee struct {
	Name  string
	Email string
}

func (e Employee) GetEmail() string {

	e.Email = "ajithmenon@gmail.com"
	return e.Email
}

func (e *Employee) GetName() string {

	e.Name = "ajith menon"
	return e.Name
}

func main() {

	a := Employee{"Ajith", "ajithm@smartrabbitz.com"}

	fmt.Printf(" Orginal Struct values \n\t name : %s \n\t email : %s \n", a.Name, a.Email)

	(&a).GetName()
	fmt.Printf("After calling Pointer receiver Method \n\t name : %s \n\t email : %s \n", a.Name, a.Email)

	a.GetEmail()
	fmt.Printf("After calling Value receiver Method \n\t name : %s \n\t email : %s \n", a.Name, a.Email)

}
