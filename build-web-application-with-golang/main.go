package main

import "fmt"

func VariadicExample(num ...int) {
	for index, item := range num {
		fmt.Printf("Index is %d and Item is %d\n", index, item);
	} 
}

func main() {
	VariadicExample(1,2,3,4,5,6,7,8,9,10)
}