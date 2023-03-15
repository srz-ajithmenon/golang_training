package main

import (
	"fmt"
)

func Hello(name string) {
	fmt.Println("Hi,", name)
}

func purchaseList(item string, price int) {
	fmt.Printf("The price of the %s is %v\n", item, price)
}

func main() {

	Hello("Ajith")
	Hello("vishnu")

	purchaseList("Laptop", 49999)
	purchaseList("Notebook", 30)

}
