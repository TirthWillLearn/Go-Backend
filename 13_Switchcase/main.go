package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// Introduction message
	fmt.Println("Switch and Case Example in Golang - Simulating a Dice Roll")

	// Seed the random number generator with the current time
	rand.Seed(time.Now().UnixNano())

	// Generate a random number between 1 and 6 (inclusive)
	diceNumber := rand.Intn(6) + 1

	// Display the dice value
	fmt.Println("Value of dice is", diceNumber)

	// Use switch-case to determine the action based on the dice number
	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 — you can open.")
	case 2:
		fmt.Println("Dice value is 2 — it's the second spot, move forward 2 steps.")
	case 3:
		fmt.Println("Dice value is 3 — third spot reached, keep going!")
	case 4:
		fmt.Println("Dice value is 4 — you're at the fourth spot, good progress.")
	case 5:
		fmt.Println("Dice value is 5 — almost there, just one more to reach the end!")
	case 6:
		fmt.Println("Dice value is 6 — jackpot! Roll the dice again.")
	default:
		fmt.Println("What was that!")
	}
}
