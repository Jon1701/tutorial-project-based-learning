// Source file belongs to the `main` package.
package main

import "fmt"

// You need an entry function called `main`.
func main() { 

	// DEFINE a variable with name `variableName` and type `type`.
	var variableName type

	// Define multiple variables `vname1`, `vname2`, and `vname3` with the same
	// type `type`.
	var vname1, vname2, vname3 type

	// Define a variable `variableName` of type `type` with initial value `value`.
	var variableName type = value;

	// Define multiple variables `vname1`, `vname2`, `vname3` all having the same
	// type `type` with initial values of `v1`, `v2`, and `v3` respectively.
	var vname1, vname2, vname3 type = v1, v2, v3

	// Can define three variables without the type and initialize their values.
	var vname1, vname2, vname3 = v1, v2, v3

	// Can remove var keyword.
	// This is called "short assignment".
	vname1, vname2, vname3 := v1, v2, v3

	// _ is a special variable. Any value given to it will be ignored.
	_, b = 34, 35
}