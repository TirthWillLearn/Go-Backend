package main

import (
	"bufio"   // For buffered reading from input
	"fmt"     // For printing to the console
	"os"      // For working with standard input (stdin)
	"strconv" // For converting strings to numbers
	"strings" // For trimming spaces and newlines
)

func main() {
	// Welcome message
	fmt.Println("Welcome to our Pizza App")
	fmt.Println("Please rate our Pizza between 1 & 5")

	// Create a buffered reader to read user input
	reader := bufio.NewReader(os.Stdin)

	// Read the input until a newline character is found
	input, _ := reader.ReadString('\n')

	// Show the raw input (includes newline)
	fmt.Println("Thanks for rating:", input)

	// Convert the input to a float after trimming whitespace/newline
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	// Handle conversion errors
	if err != nil {
		fmt.Println("Error converting rating:", err)
	} else {
		// Add 1 to the rating and display result
		fmt.Println("Added 1 to your rating:", numRating+1)
	}
}
