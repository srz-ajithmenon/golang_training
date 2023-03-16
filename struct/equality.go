package main

import "fmt"

type Data struct {
	age int
	fee int
}

func main() {

	user1 := Data{
		age: 42,
		fee: 200,
	}

	user2 := Data{
		age: 20,
		fee: 610,
	}

	if user1 == user2 {
		fmt.Println("user1 is equal to user2")
	} else {
		fmt.Println("user1 is not equal to user2")
	}

	user3 := Data{
		age: 20,
		fee: 610,
	}

	user4 := Data{
		age: 20, // fee will be 0
	}

	if user3 == user4 {
		fmt.Println("user3 is equal to user4")
	} else {
		fmt.Println("user3 is not equal to user4")
	}

	user5 := Data{
		age: 10,
		fee: 20,
	}

	user6 := Data{
		age: 10,
		fee: 20,
	}

	if user5 == user6 {
		fmt.Println("user5 is equal to user6")
	} else {
		fmt.Println("user5 is not equal to user6")
	}

}
