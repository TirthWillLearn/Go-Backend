package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in Golang")

	// Declare an array of 4 strings
	// By default, all elements are empty strings ("")
	var fruitList [4]string

	// Assign values to specific positions (indexes start from 0)
	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	// Index 2 will remain empty because we don't assign it
	fruitList[3] = "Peach"

	// Print the full array (shows all elements, including empty ones)
	fmt.Println("Fruit list is:", fruitList)

	// Print the length of the array (always fixed size once declared)
	fmt.Println("Length of fruit list:", len(fruitList))

	// Declare and initialize an array with some values
	// Remaining elements will use the zero value (empty string)
	var vegList = [5]string{"potato", "beans", "mushroom"}
	fmt.Println("Vegetable list length is:", len(vegList))
}
