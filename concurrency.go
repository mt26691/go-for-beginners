package main

import "fmt"

// concurrencyDemo is a first look at Go's concurrency. This chapter covers
// goroutines; channels and shared-state safety come in the next two chapters.
func concurrencyDemo() {
	goroutinesDemo()
}

// goroutinesDemo launches a few goroutines with `go func()`. Fill in the TODO
// (or check out the finish branch) to see them run concurrently.
func goroutinesDemo() {
	fmt.Println("\n-- Goroutines --")

	// TODO: start one goroutine per worker with `go func()`, then pause briefly
	// with time.Sleep so they finish before the function returns. (Chapter 14
	// replaces the sleep with sync.WaitGroup, the proper tool.)
	fmt.Println("  (not implemented yet)")
}
