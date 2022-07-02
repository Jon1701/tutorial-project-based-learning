package main

import "fmt"

func main() {
	// Slices are arrays of dyanmic length.
 
	// Slice of type int.
	var slice1 []int
	fmt.Printf("slice1: %s\n", slice1)

	// Define a slice and initialize its data.
	slice2 := []byte{'h','e','l','l', 'o','w','o','r','l','d'}

	// Create slices from slices.
	slice2a := slice2[3:7] // interval [x,y) = takes from indices 3,4,5,6
	fmt.Printf("slice2a: %s\n", slice2a)

	fmt.Printf("length of slice2: %d\n", len(slice2))
	fmt.Printf("maximum length of slice2: %d\n", cap(slice2))
	
	// Create a new slice.
	slice3 := []byte{'g','o','o','d','b','y','e','w','o','r','l','d'}

	// `slice4` is combined from `slice2` and `slice3`.
	slice4 := append(slice2, slice3...)
	fmt.Printf("slice4: %s\n", slice4) // helloworldgoodbyeworld

	// `slice5` contains parts of `slice2` and `slice3`.
	slice5 := append(slice2[0:5], slice3[1:4]...)
	fmt.Printf("slice5: %s\n", slice5) // helloood
}
