package main

import "fmt"

type washingMachine interface {
	wash()
	dry()
}

type samsung struct {
	kg int
}

func (w samsung) wash() {
	fmt.Println("This machine can WASH")
}

func (w samsung) dry() {
	fmt.Println("This machine can DRY")
}

func main() {

	var a washingMachine
	a = samsung{kg: 15}

	a.wash()
	a.dry()

}
