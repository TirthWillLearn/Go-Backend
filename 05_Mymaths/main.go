package main

import (
	"fmt"
	"math/big"

	// "math/rand" // Simple random number generator (not secure)
	"crypto/rand" // Cryptographically secure random number generator
)

func main() {
	fmt.Println("Welcome to maths in Go")

	// Example: Adding int and float (commented out for now)
	// var mynumberOne int = 2
	// var mynumberTwo float64 = 4.5
	// fmt.Println("The sum is: ", mynumberOne+int(mynumberTwo))

	// Example: Simple random number (commented out)
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5) + 1)

	// Generate cryptographically secure random number between 0 and 4
	myRandomNum, err := rand.Int(rand.Reader, big.NewInt(5))
	if err != nil {
		fmt.Println("Error generating secure random number:", err)
		return
	}

	fmt.Println("Crypto secure random number:", myRandomNum)
}
