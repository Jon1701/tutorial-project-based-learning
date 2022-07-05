package main

import "fmt"

func Add(a int) int {
	a = a + 1
	return a;
}

func AddByPointer(a *int) int {
	*a = *a + 1
	return *a
}

func main() {
	x := 3

	fmt.Println("========== Pass by value =========")

	fmt.Printf("x = %d\n", x)

	x1 := Add(x)

	fmt.Printf("x1 = %d\n", x1)
	fmt.Printf("x = %d\n", x)

	fmt.Println("========== Pass by reference =========")

	fmt.Printf("x = %d\n", x)

	x2 := AddByPointer(&x);

	fmt.Printf("x2 = %d\n", x2)
	fmt.Printf("x = %d\n", x)
}