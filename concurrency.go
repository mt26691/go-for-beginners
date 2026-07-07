package main

import (
	"fmt"
	"time"
)

func concurrencyDemo() {
	goroutinesDemo()
	channelsDemo()
}

// goroutinesDemo still uses time.Sleep from Chapter 13. In this chapter we
// refactor it to wait with a sync.WaitGroup — see the TODO.
func goroutinesDemo() {
	fmt.Println("\n-- Goroutines & WaitGroup --")

	const workers = 4
	// TODO: replace time.Sleep with a sync.WaitGroup. Call wg.Add(1) before each
	// goroutine, defer wg.Done() inside it, have each goroutine write into its own
	// slot in a results slice, then wg.Wait() and print the slots in order.
	for i := 0; i < workers; i++ {
		go func() {
			fmt.Printf("  worker %d finished\n", i)
		}()
	}
	time.Sleep(100 * time.Millisecond)
}

// channelsDemo will show a channel as a typed pipe between goroutines.
func channelsDemo() {
	fmt.Println("\n-- Channels --")

	// TODO: make an unbuffered channel, send a few values from a goroutine, close
	// it, and range over the channel to receive them. Then show a buffered
	// channel that accepts a few sends without a receiver waiting.
	fmt.Println("  (not implemented yet)")
}
