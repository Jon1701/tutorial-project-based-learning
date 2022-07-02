package main

func main() {
	// Keyword `iota` makes an `enum` which begins with 0.
	const (
		x = iota // x = 0
		y = iota // y = 1
		z = iota // z = 2
		w // If there is no expression after the constant's name, it uses the
		  // previous expression, which is `iota`, so w = 3
	)

	const v = iota // Once `iota` meets the keyword `const`, it resets to 0
								 // v = 0

	const (
		e,f,g = iota, iota, iota // Since all on same line, they all have value 0.
	)
}
