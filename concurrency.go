package main

import (
	"fmt"
	"sync"
)

// concurrencyDemo brings the three concurrency tools together: goroutines and a
// WaitGroup, channels, and a mutex-protected counter. Each sub-demo waits for
// the work it starts, so the output is the same on every run.
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

// channelsDemo shows a channel as a typed pipe between goroutines. The first
// example uses an unbuffered channel: a send blocks until a receiver is ready,
// so the producer and main hand values over one at a time. Ranging over the
// channel receives values until it is closed. The second example uses a buffered
// channel, whose buffer lets a few sends complete without a receiver waiting.
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
