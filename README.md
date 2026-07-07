# Go Programming for Beginners: Build Real Backend Services

This repository contains the source code for the [Go Programming for Beginners: Build Real Backend Services](https://dalabs.academy) course.

## Prerequisites

- [Go](https://go.dev/dl/) 1.22 or newer (this repo was built with Go 1.26)
- [golangci-lint](https://golangci-lint.run/) (install with `brew install golangci-lint`)
- Comfort using a terminal and a code editor (VS Code with the Go extension recommended)
- [Git](https://git-scm.com/) installed

## Start Branch

```bash
git checkout 15-mutexes-and-race-detector-start
```

The start branch carries Chapter 14's `goroutinesDemo` and `channelsDemo`, plus a stubbed `SafeCounter` (empty struct, empty `Inc`/`Value`) and a stubbed `mutexCounterDemo`. Fill in the TODOs to guard the counter with a `sync.Mutex` (or check out the finish branch).

## Finish Branch

```bash
git checkout 15-mutexes-and-race-detector-finish
```

`SafeCounter` guards an `int` with a `sync.Mutex`: `Inc` and `Value` `Lock()` before touching the value and `defer Unlock()`. `mutexCounterDemo` increments one counter from 100 goroutines at once and the final total is always exactly 100. The whole program is race-free under `go run -race .`.

## Lesson

[View the lesson on dalabs.academy](<!-- dalabs:15-mutexes-and-race-detector -->)

## Running the Program

```bash
make run     # go run .   -> runs the program
make build   # go build   -> compiles the binary
make fmt     # go fmt ./... -> formats the code
make vet     # go vet ./... -> reports suspicious code
make lint    # golangci-lint run -> runs the linter
make help    # list the available targets
```

> **Note:** This is the finish branch. A `sync.Mutex` keeps a shared `SafeCounter` correct under 100 concurrent goroutines (total always 100), and `go run -race .` reports no data races. This is the exact pattern that makes the in-memory task store safe later in the course.

## Contact

If you have any questions, feedback, or just want to connect, feel free to reach out to me on LinkedIn:

https://www.linkedin.com/in/mt26691/
