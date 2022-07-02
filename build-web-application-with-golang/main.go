package main

// Basic form
import "fmt"
import "os"

// Group form
import (
	"fmt"
	"os"
)

func main() {
 	// Basic form
	const i = 100
	const pi = 3.1415
	const prefix = "Go_"

	var i int
	var pi float
	var prefix string

	// Group form.
	const(
		i = 100
		pi = 3.1415
		prefix = "Go_"
	)

	var(
		i int
		pi float32
		prefix string
	)
}
