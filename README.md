# Go Programming for Beginners: Build Real Backend Services

This repository contains the source code for the [Go Programming for Beginners: Build Real Backend Services](https://dalabs.academy) course.

## Prerequisites

- [Go](https://go.dev/dl/) 1.22 or newer (this repo was built with Go 1.26)
- [golangci-lint](https://golangci-lint.run/) (install with `brew install golangci-lint`)
- Comfort using a terminal and a code editor (VS Code with the Go extension recommended)
- [Git](https://git-scm.com/) installed

## Start Branch

```bash
git checkout 24-testing-handlers-start
```

The start branch is the finish state of Chapter 23: the Task API is layered into `handler.go`, `service.go`, and `store.go` behind the `TaskStore` interface, wired by dependency injection in `main.go`. Everything compiles and the server runs, but there is not a single test in the project yet — `go test ./...` reports `no test files` for every package.

## Finish Branch

```bash
git checkout 24-testing-handlers-finish
```

The finish branch adds the first tests: `internal/task/handler_test.go` exercises the HTTP handlers in-process with `net/http/httptest`, injecting a fake store through the same `TaskStore` seam so no database is needed. A `make test` target runs `go test ./...`.

## Lesson

[View the lesson on dalabs.academy]({URL})
<!-- After publishing, the /publish-chapter skill replaces the placeholder above with the actual URL -->

## Project Layout

```
cmd/
  server/
    main.go        # entry point: store -> service -> handler wiring (DI), then starts the server
internal/
  task/
    task.go        # the Task resource (model) + Validate()
    store.go       # in-memory Store (map[int]Task + sync.Mutex) + compile-time TaskStore assertion
    service.go     # TaskStore interface + Service (business logic) between handlers and storage
    handler.go     # task HTTP handlers holding a *Service + writeJSON / writeError / writeServiceError helpers
```

## Running the Server

```bash
make run     # go run ./cmd/server
```

```bash
make build   # go build ./... -> compiles every package
make vet     # go vet ./...   -> reports suspicious code
make lint    # golangci-lint run -> runs the linter
```

## Running Tests

```bash
go test ./...
```

> **Note:** This is the start branch. There are no tests yet, so `go test ./...` prints `no test files` for each package. Chapter 24 adds the first handler tests with `httptest`, and a `make test` target to run them.

## Contact

If you have any questions, feedback, or just want to connect, feel free to reach out to me on LinkedIn:

https://www.linkedin.com/in/mt26691/
