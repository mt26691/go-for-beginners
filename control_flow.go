package main

import (
	"fmt"

	"github.com/mt26691/go-for-beginners/mathx"
)

// loopsDemo shows Go's single loop keyword, for: the classic three-clause
// form and the for range form over a slice.
func loopsDemo() {
	fmt.Println("\n-- Loops --")

	fmt.Print("count up: ")
	for i := 1; i <= 5; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	tasks := []string{"write code", "run tests", "ship it"}
	for i, title := range tasks {
		fmt.Printf("task %d: %s\n", i, title)
	}
}

// fizzBuzzDemo shows a condition-less switch with multiple-value cases using
// the classic FizzBuzz example. Each case is a boolean test, and Go breaks
// out of the switch automatically after the first matching case.
func fizzBuzzDemo() {
	fmt.Println("\n-- FizzBuzz --")

	for n := 1; n <= 15; n++ {
		switch {
		case n%15 == 0:
			fmt.Println("FizzBuzz")
		case n%3 == 0:
			fmt.Println("Fizz")
		case n%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(n)
		}
	}
}

// ifAndSwitchDemo shows the `if init; cond` form paired with an error and a
// switch on a value with multiple values per case.
func ifAndSwitchDemo() {
	fmt.Println("\n-- if & switch --")

	if avg, err := mathx.Average(8, 6, 10); err != nil {
		fmt.Println("average error:", err)
	} else {
		fmt.Printf("average score: %.1f\n", avg)
	}

	for _, status := range []string{"todo", "in-progress", "done", "archived"} {
		switch status {
		case "todo", "in-progress":
			fmt.Printf("%-12s -> still open\n", status)
		case "done":
			fmt.Printf("%-12s -> finished\n", status)
		default:
			fmt.Printf("%-12s -> unknown\n", status)
		}
	}
}
