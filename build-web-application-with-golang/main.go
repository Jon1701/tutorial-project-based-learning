package main

import "fmt"

func main() {
	integer := 6
	switch integer {
	case 4:
		fmt.Println("integer <= 4")
		fallthrough;
	
	case 5:
		fmt.Println("integer <= 5");
		fallthrough;

	case 6:
		fmt.Println("integer <= 6");
		fallthrough;

	case 7:
		fmt.Println("integer <= 7");
		fallthrough;

	case 8:
		fmt.Println("integer <= 8");
		fallthrough;

	default:
		fmt.Println("default case")
	}
}
