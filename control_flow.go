package main

import "fmt"

// loopsDemo will show Go's single loop keyword, for: the three-clause form
// and the for range form over a slice.
func loopsDemo() {
	fmt.Println("\n-- Loops --")

	// TODO: write a three-clause for loop (for i := 1; i <= 5; i++) that
	// prints the numbers 1 through 5.

	// TODO: range over a []string of task titles and print each index and
	// value with `for i, title := range tasks`.
}

// fizzBuzzDemo will show a condition-less switch with multiple-value cases
// using the classic FizzBuzz example.
func fizzBuzzDemo() {
	fmt.Println("\n-- FizzBuzz --")

	// TODO: loop from 1 to 15 and use a `switch {` with no condition, where
	// each case is a boolean test, to print Fizz / Buzz / FizzBuzz / the
	// number.
}

// ifAndSwitchDemo will show the `if init; cond` form paired with an error
// and a switch on a value with multiple values per case.
func ifAndSwitchDemo() {
	fmt.Println("\n-- if & switch --")

	// TODO: use `if avg, err := mathx.Average(...); err != nil` to handle the
	// error case, then print the average.

	// TODO: switch on a status string with multiple values per case
	// (for example "todo", "in-progress", "done").
}
