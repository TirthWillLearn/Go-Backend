package main

import "fmt"

// Public constant (starts with a capital letter, accessible from other packages)
const LoginToken string = "hbhvsgeyscsnvghugr"

func main() {

	// Explicitly declared string variable
	var username string = "Tirth Patel"
	fmt.Println("hello ", username, "!")
	fmt.Printf("Variable is of Type : %T \n", username)

	// Boolean variable
	var isLoggedIN bool = true
	fmt.Println(isLoggedIN)
	fmt.Printf("Variable is of Type : %T \n", isLoggedIN)

	// Unsigned 8-bit integer (0 to 255)
	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of Type : %T \n", smallVal)

	// Float32 (limited precision compared to float64)
	var smallFloat float32 = 255.455782745824839530590595
	fmt.Println(smallFloat) // Notice float32 rounds the value
	fmt.Printf("Variable is of Type : %T \n", smallFloat)

	// Default value (zero value) for int is 0
	var anotherVar int
	fmt.Println(anotherVar)
	fmt.Printf("Variable is of Type : %T \n", anotherVar)

	// Implicit type (type inferred by Go)
	var website = "Hi welcome to www.tirthde.in"
	fmt.Println(website)

	// Short variable declaration (only inside functions)
	numberofUsers := 5
	fmt.Println(numberofUsers)

	// Printing the constant value
	fmt.Println(LoginToken)
	fmt.Printf("Variable is of Type : %T \n", LoginToken)

}
