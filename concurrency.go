package main

import (
	"fmt"
	"time"
)

// concurrencyDemo is a first look at Go's concurrency. This chapter covers
// goroutines; channels and shared-state safety come in the next two chapters.
func concurrencyDemo() {
	goroutinesDemo()
}

// goroutinesDemo launches one goroutine per worker with `go func()`. The
// goroutines run concurrently, so their lines can appear in any order. We do not
// have a way to wait for goroutines yet, so we pause with time.Sleep to let them
// finish. Chapter 14 replaces that pause with sync.WaitGroup, the real tool.
func goroutinesDemo() {
	fmt.Println("\n-- Goroutines --")

	const workers = 4
	for i := 0; i < workers; i++ {
		go func() {
			fmt.Printf("  worker %d finished\n", i)
		}()
	}

	time.Sleep(100 * time.Millisecond)
}
