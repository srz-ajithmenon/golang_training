package main

import "fmt"

type Items struct {
	string
	int
}

func main() {

	item1 := Items{
		string: "bassHeads 225",
		int:    500,
	}

	fmt.Println(item1.string)
	fmt.Println(item1.int)

}

// but cannot use a type more than once
