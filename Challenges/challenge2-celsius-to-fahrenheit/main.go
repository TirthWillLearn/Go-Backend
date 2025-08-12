package main

import (
	"bufio"   // for buffered reading from input
	"fmt"     // for printing and scanning
	"os"      // for standard input (stdin)
	"strconv" // for converting strings to numbers
	"strings" // for trimming spaces/newlines
)

func main() {

	// ----------- Method 1: Using bufio.NewReader -----------

	fmt.Println("=== Celsius to Fahrenheit (bufio.NewReader) ===")
	fmt.Print("Enter temperature in Celsius: ")

	// Create a buffered reader to read from stdin
	reader := bufio.NewReader(os.Stdin)

	// Read input until newline
	input, _ := reader.ReadString('\n')

	// Remove spaces/newline
	input = strings.TrimSpace(input)

	// Convert string to float
	celsius, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println("Invalid input. Please enter a number.")
		return
	}

	// Convert Celsius to Fahrenheit
	fahrenheit := (celsius * 9 / 5) + 32

	// Print with 2 decimal places
	fmt.Printf("%.2f°C is %.2f°F\n", celsius, fahrenheit)

	// ----------- Method 2: Using fmt.Scanf -----------

	fmt.Println("\n=== Celsius to Fahrenheit (fmt.Scanf) ===")
	fmt.Print("Enter temperature in Celsius: ")

	var celsius2 float64
	// Directly scan number from input
	fmt.Scanf("%f", &celsius2)

	// Convert Celsius to Fahrenheit
	fahrenheit2 := (celsius2 * 9 / 5) + 32

	// Print result
	fmt.Printf("%.2f°C is %.2f°F\n", celsius2, fahrenheit2)
}
