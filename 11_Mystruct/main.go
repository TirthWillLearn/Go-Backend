package main

import "fmt"

// Define a struct type named 'User'
// A struct is a collection of fields (like a custom data type)
type User struct {
	Name   string // Name of the user
	Email  string // Email address
	Status bool   // Active/inactive status
	Age    int    // Age of the user
}

func main() {
	fmt.Println("Structs in Golang")
	// No inheritance in Go — no super or parent classes

	// Create a new 'User' struct instance
	Tirth := User{"Tirth Patel", "tirth@go.dev", true, 16}

	// Print the struct (will show values without field names)
	fmt.Println(Tirth)

	// Print struct with field names using %+v format specifier
	fmt.Printf("Tirth details are: %+v\n", Tirth)

	// Access individual fields from the struct
	fmt.Printf("Name is %v and email is %v.\n", Tirth.Name, Tirth.Email)
}
