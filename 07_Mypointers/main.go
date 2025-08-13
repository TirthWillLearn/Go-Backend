package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on pointers")

	// Example of a nil pointer (uncomment to see nil value)
	// var ptr *int
	// fmt.Println("Value of pointer is", ptr) // Output: <nil>

	// Declare a normal integer variable
	myNumber := 23

	// Create a pointer that stores the address of myNumber
	var ptr = &myNumber

	// Printing the pointer variable shows the memory address
	fmt.Println("Address of myNumber:", ptr)

	// Dereferencing the pointer (*) gives the value stored at that address
	fmt.Println("Value stored at pointer address:", *ptr)

	// Modify the value at the pointer address
	*ptr = *ptr + 2

	// Printing myNumber now shows the updated value
	fmt.Println("New value of myNumber:", myNumber)
}
