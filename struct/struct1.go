package main

import "fmt"

type person struct {
	first_name string
	last_name  string
	ContactInfo
}

type ContactInfo struct {
	email   string
	pincode int
}

func (p person) print() {
	fmt.Printf("%+v \n", p)
}

func (p *person) updateName(newFirstName string) {
	(*p).first_name = newFirstName
}

func main() {

	ajith := person{
		first_name: "Ajith",
		last_name:  "Menon",
		ContactInfo: ContactInfo{
			email:   "ajithm@gail.com",
			pincode: 680101,
		},
	}

	ajith.updateName("abith")
	ajith.print()
}
