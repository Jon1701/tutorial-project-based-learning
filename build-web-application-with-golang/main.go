package main

import "fmt"

func main() {
	i := 10;
	switch i {
	case 1:
		fmt.Println("i is 1");

	case 2,3,4:
		fmt.Println("i is either 2, 3, or 4");

	case 10:
		fmt.Println("i is 10");

	default:
		fmt.Println("some other integer")
	}
}
