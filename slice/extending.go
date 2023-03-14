package main

import "fmt"

func main() {

	points := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("points = %v\n", points)
	fmt.Printf("length = %v\n", len(points))

	points = append(points, 7, 8, 9) // Appending
	fmt.Printf("After appending to points = %v\n", points)
	fmt.Printf("length = %v\n", len(points))

	points = points[2:4] //Re-slice
	fmt.Printf("After Re-slicing points = %v\n", points)
	fmt.Printf("length = %v\n", len(points))

	new_points := make([]int, 3)
	copy(new_points, points)
	fmt.Printf("new_points = %v\n", points)
	fmt.Printf("length of new_points = %v\n", len(points))

}
