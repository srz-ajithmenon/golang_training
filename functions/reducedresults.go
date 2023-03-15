package main

import "fmt"

func num(x float64) (a, b, c float64) {

	a = x - 0.25
	b = x - 0.50
	c = x - 0.75

	return

}

func main() {

	m, n, p := num(10)
	fmt.Println(m, n, p)

}
