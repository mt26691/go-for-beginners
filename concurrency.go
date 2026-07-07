package main

import (
	"fmt"
	"sync"
)

func concurrencyDemo() {
	goroutinesDemo()
	channelsDemo()
}

// goroutinesDemo launches one goroutine per worker and waits for all of them
// with a sync.WaitGroup. Each goroutine writes into its own slot in results, so
// no two goroutines touch the same memory, and we print the slots in order after
// they finish for stable output.
func goroutinesDemo() {
	fmt.Println("\n-- Goroutines & WaitGroup --")

	const workers = 4
	results := make([]string, workers)

	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1) // count one more goroutine to wait for
		go func() {
			defer wg.Done() // mark this goroutine done when it returns
			results[i] = fmt.Sprintf("worker %d finished", i)
		}()
	}

	wg.Wait() // block until every wg.Done() has run

	for _, line := range results {
		fmt.Println("  " + line)
	}
}

// channelsDemo shows a channel as a typed pipe between goroutines. The first
// example uses an unbuffered channel: a send blocks until a receiver is ready,
// so the producer and main hand values over one at a time. Ranging over the
// channel receives values until it is closed. The second example uses a buffered
// channel, whose buffer lets a few sends complete without a receiver waiting.
func channelsDemo() {
	fmt.Println("\n-- Channels --")

	// Unbuffered channel: send blocks until main is ready to receive.
	nums := make(chan int)
	go func() {
		for i := 1; i <= 3; i++ {
			nums <- i // blocks until the receive below runs
		}
		close(nums) // tells the range loop there are no more values
	}()

	fmt.Print("  received from unbuffered channel:")
	for n := range nums {
		fmt.Printf(" %d", n)
	}
	fmt.Println()

	// Buffered channel: the buffer holds 3 values, so these three sends complete
	// without any receiver waiting. A fourth send would block until we received.
	letters := make(chan string, 3)
	letters <- "a"
	letters <- "b"
	letters <- "c"
	close(letters)

	fmt.Print("  received from buffered channel:")
	for s := range letters {
		fmt.Printf(" %s", s)
	}
	fmt.Println()
}
