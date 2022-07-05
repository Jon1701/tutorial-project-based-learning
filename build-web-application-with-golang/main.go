package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	width float64
	height	float64
}

type Circle struct {
	radius float64
}

// Method.
func (c Circle) Area() float64 {
	return c.radius * c.radius * math.Pi
}

// Method
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func main() {
	c1 := Circle{10}
	c2 := Circle{25}
	r1 := Rectangle{12, 2}
	r2 := Rectangle{9, 4}

	fmt.Printf("The area of r1 is: %f\n", r1.Area())
	fmt.Printf("The area of r2 is: %f\n", r2.Area())
	fmt.Printf("The area of c1 is: %f\n", c1.Area())
	fmt.Printf("The area of c2 is: %f\n", c2.Area())
	
	
}