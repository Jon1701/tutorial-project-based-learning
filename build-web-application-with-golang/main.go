package main

import (
	"errors"
)

func main() { 
	// Go has one `error` type for dealing with error messages.
	err := errors.New("this is an error message")
}