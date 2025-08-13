package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of Golang")

	// Get the current date and time
	presentTime := time.Now()
	fmt.Println(presentTime) // Prints in default format with date, time, and timezone

	// Format the current time using a reference layout
	// In Go, you must use this exact date/time (01-02-2006 15:04:05) as the layout example
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	// Create a specific date and time
	createdDate := time.Date(2020, time.August, 12, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate)

	// Format the created date to show only date and weekday
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}
