package main

import "fmt"

func main() {
	fmt.Println("== Variables, Types & Constants ==")

	variablesDemo()
	zeroValuesDemo()
	basicTypesDemo()
	constantsDemo()
}

// variablesDemo shows the two ways to declare a variable and the
// difference between declaring (:=) and assigning (=).
func variablesDemo() {
	fmt.Println("\n-- Variables --")

	// TODO: declare a variable with the long form `var`, e.g.
	//   var language string = "Go"
	// then print it.

	// TODO: declare a variable with the short form `:=`, e.g.
	//   year := 2009
	// then print it.

	// TODO: reassign one of the variables above with `=` (no :=)
	// and print it again to show declare vs assign.
}

// zeroValuesDemo shows that every variable starts at a sensible default
// when you declare it without a value.
func zeroValuesDemo() {
	fmt.Println("\n-- Zero values --")

	// TODO: declare an int, a float64, a string, and a bool with `var`
	// but no initial value, then print each one to see its zero value
	// (0, 0, "", false).
}

// basicTypesDemo shows Go's common basic types and type inference.
func basicTypesDemo() {
	fmt.Println("\n-- Basic types --")

	// TODO: declare values of a few basic types (int, int64, float64,
	// string, bool) and print each with its type using %T, e.g.
	//   fmt.Printf("%v has type %T\n", value, value)

	// TODO: show a byte and a rune from a string and print them.
}

// constantsDemo shows constants and a simple iota enumeration.
func constantsDemo() {
	fmt.Println("\n-- Constants --")

	// TODO: declare a single const, e.g. const pi = 3.14159, and print it.

	// TODO: declare a const block that uses iota to number a few values
	// (for example weekdays or log levels), then print them.
}
