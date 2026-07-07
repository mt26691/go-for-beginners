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

The start branch adds a new `concurrency.go` with `concurrencyDemo()` calling a stubbed `goroutinesDemo`, plus a new `== Concurrency ==` section in `main.go`. It compiles, vets, lints, and runs (`make run`), printing the earlier sections then a `(not implemented yet)` placeholder — ready for you to launch the goroutines.

## Finish Branch

```bash
git checkout 13-goroutines-and-concurrency-finish
```

`goroutinesDemo` launches one goroutine per worker with `go func()`, then pauses with `time.Sleep` to let them finish before returning. Because the goroutines run concurrently, their lines can print in any order. We do not have a way to *wait* for goroutines yet — Chapter 14 replaces the sleep with `sync.WaitGroup`, the proper tool.

## Lesson

[View the lesson on dalabs.academy](<!-- dalabs:13-goroutines-and-concurrency -->)

## Running the Program

```bash
make run     # go run .   -> runs the program
make build   # go build   -> compiles the binary
make fmt     # go fmt ./... -> formats the code
make vet     # go vet ./... -> reports suspicious code
make lint    # golangci-lint run -> runs the linter
make help    # list the available targets
```

> **Note:** This chapter introduces goroutines only. Waiting is done with a short `time.Sleep`; Chapter 14 replaces it with `sync.WaitGroup` and adds channels.

## Contact

If you have any questions, feedback, or just want to connect, feel free to reach out to me on LinkedIn:

https://www.linkedin.com/in/mt26691/
