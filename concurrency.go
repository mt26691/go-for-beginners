package main

import (
	"fmt"
	"sync"
)

// concurrencyDemo is a gentle first look at Go's concurrency tools: goroutines,
// channels, and the sync package. Each sub-demo waits for the work it starts and
// prints its results in a fixed order, so the output is the same on every run.
func concurrencyDemo() {
	goroutinesDemo()
	channelsDemo()
	mutexCounterDemo()
}

// goroutinesDemo launches a few goroutines with `go func()` and waits for all of
// them to finish with a sync.WaitGroup before printing. The goroutines run in an
// unspecified order, so each one writes into its own slot in results — no shared
// write collides — and we print the slots in order afterwards for stable output.
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

// SafeCounter is a counter guarded by a sync.Mutex so it is safe to use from
// many goroutines at once. Lock before touching the shared value and Unlock
// after, so only one goroutine reads or writes it at a time. This is the same
// pattern that makes the in-memory task store safe under concurrent HTTP
// requests later in the course.
type SafeCounter struct {
	mu    sync.Mutex
	value int
}

// Inc adds one to the counter while holding the lock.
func (c *SafeCounter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Value returns the current count while holding the lock.
func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

// mutexCounterDemo hammers a single SafeCounter from many goroutines at once.
// Because every increment goes through the mutex, the final total is exact and
// deterministic: 100 goroutines each adding one always end at 100, with no data
// race. Without the mutex this would race and the total would vary run to run.
func mutexCounterDemo() {
	fmt.Println("\n-- Mutex-protected counter --")

	const goroutines = 100
	counter := &SafeCounter{}

	var wg sync.WaitGroup
	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Inc()
		}()
	}

	wg.Wait()

	fmt.Printf("  %d goroutines each +1 -> final total: %d\n", goroutines, counter.Value())
}
