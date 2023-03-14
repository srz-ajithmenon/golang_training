package main

import "fmt"

func main() {

	z := make(map[string]string)
	z["name"] = "Ashique"
	z["school"] = "SNGIST"
	z["class"] = "Final Year"
	z["Grade"] = "4.5"
	z["email"] = "ashique@gmail.com"

	fmt.Println("Map : ", z)
	fmt.Println("School name is ", z["school"])

	z["place"] = "kodungalloor"
	fmt.Println("After appending : ", z)

	z["name"] = "Amalraj"
	fmt.Println("After changing name : ", z)

	delete(z, "class")
	fmt.Println("After deleting class : ", z)

	// checking existance

	val1, key1 := z["name"]
	val2, key2 := z["age"]

	fmt.Println("name: ", val1, key1)
	fmt.Println("age: ", val2, key2)

}
