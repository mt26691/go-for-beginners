package main

import "fmt"

// concurrencyDemo is a gentle first look at Go's concurrency tools: goroutines,
// channels, and the sync package. The three sub-demos are stubbed on this start
// branch — fill in the TODOs (or check out the finish branch) to see them work.
func concurrencyDemo() {
	goroutinesDemo()
	channelsDemo()
	mutexCounterDemo()
}

// goroutinesDemo launches a few goroutines with `go func()` and waits for all of
// them to finish with a sync.WaitGroup before printing a sorted summary. Sorting
// the results keeps the output stable even though the goroutines run in an
// unspecified order.
func goroutinesDemo() {
	fmt.Println("\n-- Goroutines & WaitGroup --")

	// TODO: launch a goroutine per worker with `go func()`, track them with a
	// sync.WaitGroup, Wait() for all of them, then print the collected results
	// in sorted order so the output is deterministic.
	fmt.Println("  (not implemented yet)")
}

// channelsDemo shows a channel as a typed pipe between goroutines: one goroutine
// sends values, main receives them. It contrasts an unbuffered channel (send
// blocks until a receiver is ready) with a buffered one (send does not block
// until the buffer is full).
func channelsDemo() {
	fmt.Println("\n-- Channels --")

	// TODO: make an unbuffered channel, send a known number of values from a
	// goroutine, close it, and range over the channel to receive them. Then show
	// a buffered channel that accepts a few sends without a receiver.
	fmt.Println("  (not implemented yet)")
}

// SafeCounter is a counter guarded by a sync.Mutex so it is safe to use from
// many goroutines at once. Lock() before touching the shared value and Unlock()
// after, so only one goroutine reads or writes it at a time. This is the same
// pattern that makes the in-memory task store safe under concurrent HTTP
// requests later in the course.
type SafeCounter struct {
	// TODO: add a sync.Mutex and an int value field.
}

// Inc adds one to the counter while holding the lock.
func (c *SafeCounter) Inc() {
	// TODO: lock, increment the value, unlock (use defer for the unlock).
}

// Value returns the current count while holding the lock.
func (c *SafeCounter) Value() int {
	// TODO: lock, read the value, unlock, and return it.
	return 0
}

// mutexCounterDemo hammers a single SafeCounter from many goroutines at once.
// Because every increment goes through the mutex, the final total is exact and
// deterministic: N goroutines each adding one always end at N, with no data race.
func mutexCounterDemo() {
	fmt.Println("\n-- Mutex-protected counter --")

	// TODO: launch N goroutines that each call counter.Inc(), wait for them all
	// with a WaitGroup, then print counter.Value() — which is always N.
	fmt.Println("  (not implemented yet)")
}
