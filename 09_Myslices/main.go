package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to video on slices")

	// Create a slice of strings (fruitList)
	// Slices are dynamic arrays in Go
	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitList is %T\n", fruitList)

	// Append more fruits to the slice
	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList) // ["Apple", "Tomato", "Peach", "Mango", "Banana"]

	// Slice operation: keep only the first 3 elements
	fruitList = append(fruitList[:3])
	fmt.Println(fruitList) // ["Apple", "Tomato", "Peach"]

	// Create a slice of integers with length 4 (default values are 0)
	highScores := make([]int, 4)

	// Assign values to each index
	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
	// highScores[4] = 777 // ❌ This would cause an error because length is only 4 initially

	// Append more integers to highScores (length increases dynamically)
	highScores = append(highScores, 555, 666, 321)
	fmt.Println(highScores) // [234 945 465 867 555 666 321]

	// Sort the integers in ascending order
	sort.Ints(highScores)
	fmt.Println("Sorted highScores:", highScores)

	// Example: removing a value from a slice by index
	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println("Original courses:", courses)

	var index int = 2 // We want to remove the 3rd element ("swift")
	// Take elements before index and after index, join them together
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("After removal:", courses) // "swift" is gone
}
