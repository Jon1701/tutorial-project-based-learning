package main

import "fmt"

// We use the type `bool` to define a variable as a boolean type.
// It can only be `true` or `false`.
// The default value will be `false`.
var isActive bool; // global variable.

func test() {
	var available bool // local variable.
	valid := false // Shorthand initialization.
	available = true // Assignment of value to variable.
}

func main() { 
	fmt.Println("Hello World")
}