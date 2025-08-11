// package main

// import "fmt"

// func main() {
// 	var name string = "Tirth Patel"
// 	var experience int = 2
// 	stack := []string{"Node.js", "Go", "MySQL"}

// 	fmt.Println("Name:", name)
// 	fmt.Println("Experience:", experience, "years")
// 	fmt.Println("Stack:", stack)
// }

package main

import "fmt"

func main() {
	// Step 1: declare city with :=
	city := "YourCity"

	// Step 2: declare population with var and type int
	var population int = 100000

	// Step 3: print city and population
	fmt.Println("City:", city)
	fmt.Println("Population:", population)

	// Step 4: update population
	population = population + 5000
	fmt.Println("Updated Population:", population)

	// Call add function
	sum := add(10, 20)
	fmt.Println("Sum:", sum)

	fmt.Println("Hi I am Tirth Welcome to my Go Page")
}

// add function with typed parameters and return type
func add(a, b int) int {
	return a + b
}
