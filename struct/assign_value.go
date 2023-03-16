package main

import (
	"fmt"
)

type User1 struct {
	firstName string

	lastName string

	age int

	fee int
}

func main() {

	user1 := User1{
		age:       24,
		firstName: "Ajjith",
		fee:       10000,
		lastName:  "Menon",
	}

	user2 := User1{"Vishnu", "K Rajan", 26, 10000} // you have to enter in order of declaration

	var user3 User1

	user3.firstName = "Robish"
	user3.lastName = "A.K"
	user3.fee = 25000
	user3.age = 25

	fmt.Println("User 1:", user1)
	fmt.Println("User 2:", user2)
	fmt.Println("User 3:", user3)

}
