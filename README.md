# Go Programming for Beginners: Build Real Backend Services

This repository contains the source code for the [Go Programming for Beginners: Build Real Backend Services](https://dalabs.academy) course.

## Prerequisites

- [Go](https://go.dev/dl/) 1.22 or newer (this repo was built with Go 1.26)
- [golangci-lint](https://golangci-lint.run/) (install with `brew install golangci-lint`)
- Comfort using a terminal and a code editor (VS Code with the Go extension recommended)
- [Git](https://git-scm.com/) installed

## Start Branch

```bash
git checkout 14-channels-and-waitgroup-start
```

The start branch carries Chapter 13's `goroutinesDemo` (still using `time.Sleep`) and a stubbed `channelsDemo`. Fill in the TODOs to refactor `goroutinesDemo` onto a `sync.WaitGroup` and to implement the channel examples (or check out the finish branch).

## Finish Branch

```bash
git checkout 14-channels-and-waitgroup-finish
```

`goroutinesDemo` now waits with a `sync.WaitGroup`: `wg.Add(1)` before each goroutine, `defer wg.Done()` inside, and `wg.Wait()` at the end. Each goroutine writes into its own slot in a `results` slice, so the output is deterministic. `channelsDemo` sends values over an unbuffered channel (a send blocks until a receiver is ready) and then a buffered channel (sends succeed until the buffer is full).

## Lesson

[View the lesson on dalabs.academy](https://dalabs.academy/courses/go-programming-for-beginners-build-real-backend-services/go-language-foundations/channels-and-waitgroup)

## Running the Program

```bash
make run     # go run .   -> runs the program
make build   # go build   -> compiles the binary
make fmt     # go fmt ./... -> formats the code
make vet     # go vet ./... -> reports suspicious code
make lint    # golangci-lint run -> runs the linter
make help    # list the available targets
```

> **Note:** This chapter adds the two coordination tools — `sync.WaitGroup` (wait for goroutines) and channels (pass values between them). Shared-state safety with `sync.Mutex` comes in Chapter 15.

## Contact

If you have any questions, feedback, or just want to connect, feel free to reach out to me on LinkedIn:

https://www.linkedin.com/in/mt26691/
