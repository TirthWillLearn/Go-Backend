package main

import "fmt"

func main() {
	fmt.Println("Maps in Go (Golang)")

	// Create a map where both keys and values are strings
	// Keys will be short forms, values will be full names.
	languages := make(map[string]string)

	// Add default values to the map
	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"
	languages["MY"] = "MySQL"    // MY for MySQL
	languages["ND"] = "Node"     // ND for Node
	languages["PG"] = "Postgres" // PG for Postgres

	// Print the complete map
	fmt.Println("List of all languages:", languages)

	// Access and print a specific value using its key
	fmt.Println("JS stands for:", languages["JS"])

	// Delete a key-value pair from the map
	delete(languages, "RB") // Removing Ruby
	fmt.Println("List of all languages after deleting RB:", languages)

	// Loop through the map and print each value
	// The blank identifier "_" ignores the key in this loop
	for key, value := range languages {
		fmt.Printf("Short form: %v, Full name: %v\n", key, value)
	}
}
