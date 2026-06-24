# Go Programming for Beginners: Build Real Backend Services

This repository contains the source code for the [Go Programming for Beginners: Build Real Backend Services](https://dalabs.academy) course.

## Prerequisites

- [Go](https://go.dev/dl/) 1.22 or newer (this repo was built with Go 1.26)
- [golangci-lint](https://golangci-lint.run/) (install with `brew install golangci-lint`)
- Comfort using a terminal and a code editor (VS Code with the Go extension recommended)
- [Git](https://git-scm.com/) installed

## Start Branch

```bash
git checkout 13-goroutines-and-concurrency-start
```

The start branch sets the stage for Go's concurrency tools. A new `concurrency.go` declares `concurrencyDemo()` and three sub-demos — `goroutinesDemo`, `channelsDemo`, and `mutexCounterDemo` — plus a `SafeCounter` type whose `Inc`/`Value` methods and counter field are stubbed with `// TODO` comments. `main.go` adds a new `== Concurrency ==` section that calls `concurrencyDemo`. It compiles, vets, lints, and runs (`make run`) — printing the earlier sections, then the `Concurrency` header with `(not implemented yet)` placeholder lines — ready for you to add the `go func()` + `sync.WaitGroup` launch, the channel send/receive, and the `sync.Mutex`-protected counter.

## Finish Branch

```bash
git checkout 13-goroutines-and-concurrency-finish
```

The finish branch completes the concurrency demo. `goroutinesDemo` launches one goroutine per worker with `go func()`, waits for all of them with a `sync.WaitGroup`, then prints a fixed-order summary (each goroutine writes into its own slice slot) so the output is stable across runs. `channelsDemo` sends a known number of values over an unbuffered channel and ranges over it to receive them, then shows a buffered channel accepting sends without a receiver. `SafeCounter` guards an `int` with a `sync.Mutex`, and `mutexCounterDemo` increments one counter from 100 goroutines at once — the final total is always exactly 100, with no data race. The whole program is race-free under `go run -race .`, and `make run` prints real, deterministic output.

## Lesson

[View the lesson on dalabs.academy](https://dalabs.academy/courses/go-programming-for-beginners-build-real-backend-services/go-language-foundations/goroutines-and-concurrency)

## Running the Program

```bash
make run     # go run .   -> runs the program
make build   # go build   -> compiles the binary
make fmt     # go fmt ./... -> formats the code
make vet     # go vet ./... -> reports suspicious code
make lint    # golangci-lint run -> runs the linter
make help    # list the available targets
```

To check the program for data races, run it with Go's built-in race detector:

```bash
go run -race .
```

> **Note:** This is the finish branch, so the `== Concurrency ==` section is fully implemented. The goroutines launch and join through a `sync.WaitGroup`, values flow over an unbuffered and then a buffered channel, and a `sync.Mutex` keeps a shared `SafeCounter` correct under 100 concurrent goroutines — the final total is always exactly 100. The output is deterministic across runs, and `go run -race .` reports no data races.

## Contact

If you have any questions, feedback, or just want to connect, feel free to reach out to me on LinkedIn:

https://www.linkedin.com/in/mt26691/
