package main

import "fmt"

func main() {
	// If-else statement.
	if x > 10 {
		fmt.Println("x is greater than 10")
	} else {
		fmt.Println("x is less than or equal to 10")
	}

	// Initialize x, then check
	if x:= computedValue(); x > 10 {
		fmt.Println("x is greater than 10")
	} else {
		fmt.Println("x is less than 10")
	}
}
