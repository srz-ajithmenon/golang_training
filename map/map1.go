package main

import (
	"fmt"
)

func printMap(m map[string]string) {
	for key, val := range m {
		fmt.Println(key, " is ", val)
	}
	fmt.Println("")
}

func main() {

	var places map[string]string

	colors := map[string]string{
		"red": "#ff0000", "green": "#00ff00", "white": "#ffffff",
	}

	vehicle := make(map[string]string)

	vehicle["Swift"] = "car"
	vehicle["R3"] = "bike"
	vehicle["volvo"] = "bus"

	// places["India"] = "New Delhi"

	fmt.Println("Places \t", places)
	fmt.Println("Vehicle\t", vehicle, "\n")

	fmt.Println("Colors Before \t", colors)
	delete(colors, "red")
	fmt.Println("Colors After \t", colors, "\n")

	printMap(colors)
	printMap(vehicle)

}
