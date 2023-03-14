package main

import "fmt"

func main() {

	var x = map[int]string{1: "Ajith", 2: "Vishnu", 3: "deepthi"}

	y := map[string]int{"First": 1, "Second": 2, "third": 3, "Fourth": 4, "Fifth": 5}

	z := make(map[string]string)
	z["name"] = "Ashique"
	z["school"] = "SNGIST"
	z["class"] = "Final Year"
	z["Grade"] = "4.5"
	z["email"] = "ashique@gmail.com"

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	// initial values with different methods

	var a = make(map[string]string)
	var b = map[string]string{}
	var c map[string]string

	fmt.Println("a: ", a)
	fmt.Println(a != nil)
	fmt.Println("b: ", b)
	fmt.Println(b != nil)
	fmt.Println("c: ", c)
	fmt.Println(c != nil)

}
