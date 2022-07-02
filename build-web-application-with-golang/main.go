package main

func main() {
	// Maps are similar to dictionaries in Python, or JS objects, but with the
	// limitation that all keys must have the same type, and all values must have
	// the same type.

	// Declare a map where the keys are of type `string` and values of type `int`.
	var numbers map[string]int

	// Use `make` function to initialize it.
	num2 := make(map[string]int)

	num2["one"] = 1 // Assign values to a key.
	num2["ten"] = 10
	num2["three"] = 3

	// Order is ambiguous.
}
