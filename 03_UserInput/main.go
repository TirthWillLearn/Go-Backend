package main

import (
	"bufio" // Provides buffered input/output
	"fmt"   // For formatted input/output
	"os"    // For OS-level operations (e.g., reading from stdin)
)

func main() {

	// Welcome message
	welcome := "Welcome to User Input"
	fmt.Println(welcome)

	// Create a buffered reader to read from standard input (keyboard)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza: ")

	// Read input until a newline character ('\n') is encountered
	// The second return value is an error, which we ignore here using "_"
	input, _ := reader.ReadString('\n')

	// Output the received input
	fmt.Println("Thanks for rating,", input)

	// Show the data type of 'input' (it will always be a string from ReadString)
	fmt.Printf("Type of this rating is %T \n", input)
}
