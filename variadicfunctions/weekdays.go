package main

import "fmt"

func daysOfWeek(days ...string) {

	for i, d := range days {

		fmt.Printf("Days %d: %s\n", i+1, d)

	}

}

func main() {

	fmt.Println("__________________________")

	daysOfWeek() // will be empty

	fmt.Println("__________________________")

	daysOfWeek("Monday")

	fmt.Println("__________________________")

	daysOfWeek("Monday", "Tuesday")

	fmt.Println("__________________________")

	daysOfWeek("Monday", "Tuesday", "Wedsday", "Thursday", "Friday")

}
