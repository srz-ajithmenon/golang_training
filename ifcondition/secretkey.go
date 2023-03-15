package main

import "fmt"

func main() {

	secret_key := 102101

	if secret_key == 102101 {
		fmt.Println("Access Allowed")
	} else {
		fmt.Println("Access Denied")
	}

}
