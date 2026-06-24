package main

import (
	"fmt"

	"github.com/mt26691/go-for-beginners/mathx"
	"github.com/mt26691/go-for-beginners/textx"
)

func main() {
	fmt.Println("== Functions & Packages ==")

	functionsDemo()
	packagesDemo()
}

// functionsDemo shows functions with multiple return values, named
// returns, and a variadic parameter.
func functionsDemo() {
	fmt.Println("\n-- Functions --")

	min, max := mathx.MinMax(7, 2, 9, 4, 1)
	fmt.Printf("MinMax(7, 2, 9, 4, 1) = %d, %d\n", min, max)

	avg, err := mathx.Average(7, 2, 9, 4, 1)
	if err != nil {
		fmt.Println("Average error:", err)
	} else {
		fmt.Printf("Average(7, 2, 9, 4, 1) = %.2f\n", avg)
	}

	_, err = mathx.Average()
	if err != nil {
		fmt.Println("Average() error:", err)
	}
}

// packagesDemo imports our own textx package by its full module path and
// calls an exported function from it.
func packagesDemo() {
	fmt.Println("\n-- Packages --")

	raw := "  go    backend   service  "
	fmt.Printf("Label(%q) = %q\n", raw, textx.Label(raw))
}
