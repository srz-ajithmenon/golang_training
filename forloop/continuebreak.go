package main

import "fmt"

func main() {

	for i := 0; i < 20; i++ {

		if i%2 == 0 {
			continue
		}

		if i == 15 {
			break
		}

		fmt.Println(i)
	}

}
