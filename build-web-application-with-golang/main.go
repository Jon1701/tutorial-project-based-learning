package main

import "fmt"

type Person struct {
	name string
	age int
}

func isOlder(p1 Person, p2 Person) (Person, int) {
	if (p1.age > p2.age) {
		return p1, p1.age - p2.age
	}

	return p2, p2.age - p1.age;
}

func main() {
	var p Person // p is of type Person.

	p.name = "Astaxie"
	p.age = 25

	fmt.Printf("The person's name is: %s\n", p.name)
	fmt.Printf("The person's age is: %d\n", p.age)

	var tom Person
	tom.name = "Tom"
	tom.age = 18
	bob := Person{age: 25, name: "Bob"}
	paul := Person{"Paul", 43}

	tb_Older, tb_diff := isOlder(tom, bob);
	tp_Older, tp_diff := isOlder(tom, paul)
	bp_Older, bp_diff := isOlder(bob, paul)

	fmt.Printf("Between %s and %s, %s is older by %d years\n", tom.name, bob.name, tb_Older.name, tb_diff);
	fmt.Printf("Between %s and %s, %s is older by %d years\n", tom.name, paul.name, tp_Older.name, tp_diff);
	fmt.Printf("Between %s and %s, %s is older by %d years\n", bob.name, paul.name, bp_Older.name, bp_diff);
	
}