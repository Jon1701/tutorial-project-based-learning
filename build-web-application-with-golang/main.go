package main

import "fmt"

type Human struct {
	name string
	age int
	weight int
}

type Student struct {
	Human // embedded field
	specialty string
}

func main() {
	mark := Student{Human{"Mark", 25, 120}, "Computer Science"}

	fmt.Printf("Name is: %s\n", mark.name);
	fmt.Printf("Age is: %d\n", mark.age);
	fmt.Printf("Weight is: %d\n", mark.weight);
	fmt.Printf("Specialty is: %s\n", mark.specialty)

	mark.specialty = "AI"
	fmt.Println("Mark changed his specialty")
	fmt.Printf("Specialty is: %s\n", mark.specialty)

}