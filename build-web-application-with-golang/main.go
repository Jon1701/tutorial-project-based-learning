package main

import "fmt"

func main() { 
	// Go uses the UTF-8 character set.

	var frenchHello string

	var emptyString string = ""

	no, yes, maybe := "no", "yes", "maybe"

	japaneseHello := "ohayou"

	frenchHello := "bonjour"

	// Strings are immutable.
	// Cannot change them after creation.
	var s string = "hello"
	s[0] = "c" // compile error

	// Convert strings to []byte, modify, then convert back to string.
	c := []byte(s) 	// convert to []byte 
	c[0] = 'c' 			// modify
	s2 := string(c) // convert to string

	fmt.Printf("%s\n", a) // Print string to console.

	// Combine two strings.
	a := "hello,"
	b := " world"
	v := a + b
	fmt.Printf("%s\n", v)
}