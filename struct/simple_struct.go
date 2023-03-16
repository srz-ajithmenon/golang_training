package main

import "fmt"

type Item struct {
	no    int
	name  string
	price float64
	store string
}

func main() {

	var itm1 Item
	var itm2 Item

	itm1.no = 1
	itm1.name = "Airdopes 131 pro"
	itm1.price = 999.99
	itm1.store = "Flipkart"

	itm2.no = 2
	itm2.name = "Mi powerbank"
	itm2.price = 1000.00
	itm2.store = "MyG"

	fmt.Println("Item Number:", itm1.no)
	fmt.Println("Item Name:", itm1.name)
	fmt.Println("Price:", itm1.price)
	fmt.Println("Store:", itm1.store)

	fmt.Print("\n")

	fmt.Println("Item Number:", itm2.no)
	fmt.Println("Item Name:", itm2.name)
	fmt.Println("Price:", itm2.price)
	fmt.Println("Store:", itm2.store)

	// fmt.Print("\n")

	// fmt.Println(itm1)
	// fmt.Println(itm2)

}
