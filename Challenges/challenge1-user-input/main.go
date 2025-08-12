package main

import (
	"bufio"   // For buffered reading from input
	"fmt"     // For printing output to console
	"os"      // For accessing system's standard input/output
	"strings" // For string manipulation (like trimming spaces/newlines)
)

func main() {

	// Print a welcome message asking for user input
	fmt.Println("Welcome! Please enter your name:")

	// Create a new reader to read from the console (stdin)
	reader := bufio.NewReader(os.Stdin)

	// Read user input until they press Enter
	name, _ := reader.ReadString('\n')

	// Remove extra spaces or newline characters from the input
	name = strings.TrimSpace(name)

	// Print a greeting with the entered name
	fmt.Printf("Hello, %s !\n", name)
}
