package main

import "fmt"

func main() {
	fmt.Println("== Functions & Packages ==")

	functionsDemo()
	packagesDemo()
}

// functionsDemo will show functions with multiple return values, named
// returns, and variadic parameters.
func functionsDemo() {
	fmt.Println("\n-- Functions --")

	// TODO: call mathx.MinMax with several numbers and print both
	// returned values.

	// TODO: call mathx.Average and handle the (value, error) pair with
	// `if err != nil`.
}

// packagesDemo will import our own package by its full module path and
// call an exported function from it.
func packagesDemo() {
	fmt.Println("\n-- Packages --")

	// TODO: import github.com/mt26691/go-for-beginners/textx and call
	// textx.Label to format a raw string into a tidy display label.
}
