package main

import "fmt"

// slicesDemo will show how slices grow with append, how len and cap differ, a
// slice expression, ranging over a slice, and the aliasing gotcha that happens
// when two slices share the same backing array.
func slicesDemo() {
	fmt.Println("\n-- Slices --")

	// TODO: build a []Task with make + append, print len and cap, take a slice
	// expression, range over it, and show the shared-backing-array aliasing trap.
}

// mapsDemo will show a map[int]Task keyed by ID: set and get values, the
// comma-ok lookup for present and absent keys, delete, and why we iterate by
// sorted keys (map iteration order is unspecified).
func mapsDemo() {
	fmt.Println("\n-- Maps --")

	// TODO: build a map[int]Task, get one task, show the comma-ok idiom for a
	// present and an absent key, delete a key, and iterate in sorted-key order.
}

// nilMapDemo will show that reading from a nil map returns the zero value
// safely, while writing to a nil map panics.
func nilMapDemo() {
	fmt.Println("\n-- Nil maps --")

	// TODO: declare a nil map[int]Task, read a missing key (safe, zero value),
	// and explain in the output that writing to it would panic.
}
