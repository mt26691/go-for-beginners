package main

import (
	"fmt"
	"sync"
)

func concurrencyDemo() {
	goroutinesDemo()
	channelsDemo()
	mutexCounterDemo()
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
		wg.Add(1)
		go func() {
			defer wg.Done()
			results[i] = fmt.Sprintf("worker %d finished", i)
		}()
	}

	wg.Wait()

	for _, line := range results {
		fmt.Println("  " + line)
	}
}

// channelsDemo shows a channel as a typed pipe between goroutines: an unbuffered
// channel hands values over one at a time, and a buffered channel accepts a few
// sends without a receiver waiting.
func channelsDemo() {
	fmt.Println("\n-- Channels --")

	nums := make(chan int)
	go func() {
		for i := 1; i <= 3; i++ {
			nums <- i
		}
		close(nums)
	}()

	fmt.Print("  received from unbuffered channel:")
	for n := range nums {
		fmt.Printf(" %d", n)
	}
	fmt.Println()

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

// SafeCounter will guard an int with a sync.Mutex so it is safe to use from many
// goroutines at once. Fill in the TODOs (or check out the finish branch).
type SafeCounter struct {
	// TODO: add a sync.Mutex field (mu) and an int value field.
}

// Inc adds one to the counter while holding the lock.
func (c *SafeCounter) Inc() {
	// TODO: Lock, increment the value, Unlock (use defer for the Unlock).
}

// Value returns the current count while holding the lock.
func (c *SafeCounter) Value() int {
	// TODO: Lock, read the value, Unlock, and return it.
	return 0
}

// mutexCounterDemo will hammer a single SafeCounter from many goroutines at once.
func mutexCounterDemo() {
	fmt.Println("\n-- Mutex-protected counter --")

	// TODO: launch 100 goroutines that each call counter.Inc(), wait for them all
	// with a sync.WaitGroup, then print counter.Value() — always exactly 100.
	fmt.Println("  (not implemented yet)")
}
