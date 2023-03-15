package main

import "fmt"

func main() {
	age := 10

	if age < 18 {
		fmt.Println("Your age is below the required age")
	} else if age >= 18 && age <= 65 {
		fmt.Println("you are eligible")
	} else if age > 65 && age < 70 {
		fmt.Println("Under Consideration")
	} else {
		fmt.Println("you are too old to apply")
	}
}
