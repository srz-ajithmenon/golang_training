package main

import (
	"fmt"
)

func main() {

	var a1 = [4]int{7, 8, 9, 10}
	a2 := [6]int{11, 12, 13, 14, 15, 16}
	a3 := [...]int{2, 3, 4, 5} // ... will allow go to decide array length

	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)

}
