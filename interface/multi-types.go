package main

import "fmt"

type Plant interface {
	Grow() string
}

type Mango struct {
}

func (m Mango) Grow() string {
	return "The mango tree is growing"
}

type Apple struct {
}

func (a Apple) Grow() string {
	return "The apple tree is growing!"
}

type Grapes struct {
}

func (g Grapes) Grow() string {
	return "The grape plant is growing!"
}

type Any_type struct {
}

func (j Any_type) Grow() string {
	return "You can use any type with the method"
}

func main() {

	plants := []Plant{Mango{}, Apple{}, Grapes{}, Any_type{}}

	for _, v := range plants {
		fmt.Println(v.Grow())
	}

}
