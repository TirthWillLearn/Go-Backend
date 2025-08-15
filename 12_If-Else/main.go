package main

import "fmt"

func main() {
	fmt.Println("If else in Golang")

	loginCount := 10
	var result string

	// Basic if-else-if statement
	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount > 10 {
		result = "Watch out"
	} else {
		result = "Exactly 10 login count"
	}

	fmt.Println(result)

	// Check if a number is even or odd
	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	// Short statement in if: declare 'num' and check its value
	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is NOT less than 10")
	}

	// Common Go pattern for error checking
	// if err != nil {
	//     // handle error
	// }
}
