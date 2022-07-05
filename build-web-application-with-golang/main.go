package main

import "fmt"

func SumAndProduct(a int, b int) (int, int) {
	return a+b, a*b
}

func main() {
	x := 3;
	y := 4;

	xPLUSy, xTIMESy := SumAndProduct(x, y);

	fmt.Printf("%d + %d = %d\n", x, y, xPLUSy);
	fmt.Printf("%d * %d = %d\n", x, y, xTIMESy);
}