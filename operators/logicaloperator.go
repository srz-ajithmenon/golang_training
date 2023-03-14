package main

import "fmt"

func main() {

	x := 5
	y := 10
	z := 15

	fmt.Println(" x is less than y and x is greater than z : ", x < y && x > z)
	fmt.Println(" x is less than y or x is greater than z : ", x < y || x > z)
	fmt.Println(" x is not less than y and x is not greater than z : ", !(x < y) && !(x > z))

}
