package main

import "fmt"

func main() {

	month := "February"

	switch month {

	case "January", "March", "May", "July", "August", "October", "December":

		fmt.Printf("There are 31 days in %s\n", month)

	case "April", "June", "September", "November":

		fmt.Printf("There are 30 days in %s\n", month)

	case "February":

		fmt.Printf("There are 28 days in %s\n", month)

	default:

		fmt.Printf("Invalid name of month\n", month)

	}

}
