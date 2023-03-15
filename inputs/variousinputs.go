package main

import "fmt"

func main() {

	var name1, name2, name3 string

	var age1, age2, age3 int

	fmt.Println("(scanln) Enter Name and Age: ")
	fmt.Scanln(&name1, &age1) // will only allow space to scan value   // if pressed enter in b/w it execute scanning
	fmt.Printf("%s's age is %d\n", name1, age1)

	fmt.Println("(scan) Enter Name and Age: ")
	fmt.Scan(&name2, &age2) // will allow both space and enter to scan value
	fmt.Printf("%s's age is %d\n", name2, age2)

	// fmt.Println("(scanf-' ') Enter Name and Age: ")
	// fmt.Scanf("%v %v", &name3, &age3) // will only allow space to scan value   // if pressed enter in b/w it execute scanning
	// fmt.Printf("%s's age is %d\n", name3, age3)

	fmt.Println("(scanf-' ') Enter Name and Age: ")
	fmt.Scanf("%v\n%v", &name3, &age3) // will only allow enter to scan value   // if pressed space in b/w it execute scanning
	fmt.Printf("%s's age is %d\n", name3, age3)

}
