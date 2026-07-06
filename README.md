# Go Programming for Beginners: Build Real Backend Services

This repository contains the source code for the [Go Programming for Beginners: Build Real Backend Services](https://dalabs.academy) course.

## Prerequisites

- [Go](https://go.dev/dl/) 1.22 or newer (this repo was built with Go 1.26)
- [golangci-lint](https://golangci-lint.run/) (install with `brew install golangci-lint`)
- Comfort using a terminal and a code editor (VS Code with the Go extension recommended)
- [Git](https://git-scm.com/) installed

## Start Branch

```bash
git checkout 25-table-driven-tests-start
```

The start branch is the finish state of Chapter 24: `internal/task/handler_test.go` tests the HTTP handlers in-process with `net/http/httptest` and a `fakeStore`. The tests are written as **individual functions and `t.Run` subtests** — `TestCreateTask` (valid → 201, empty title → 400), `TestListTasks` (→ 200 array), and `TestGetTask` (found → 200, missing → 404). They pass, but the create/get cases repeat the same request/record/assert shape, and they do not yet cover a too-long title or a non-numeric id. This is the "before" that Chapter 25 refactors.

## Finish Branch

```bash
git checkout 25-table-driven-tests-finish
```

The finish branch refactors those repetitive tests into the idiomatic **table-driven** form: a slice of case structs iterated with `t.Run(tc.name, ...)`, marked `t.Parallel()`, covering **more** cases in **fewer**, denser functions.

## Lesson

[View the lesson on dalabs.academy](<!-- dalabs:25-table-driven-tests -->)
<!-- After publishing, the /publish-chapter skill replaces the placeholder above with the actual URL -->

## Project Layout

```
cmd/
  server/
    main.go        # entry point: store -> service -> handler wiring (DI), then starts the server
internal/
  task/
    task.go          # the Task resource (model) + Validate()
    store.go         # in-memory Store (map[int]Task + sync.Mutex) + compile-time TaskStore assertion
    service.go       # TaskStore interface + Service (business logic) between handlers and storage
    handler.go       # task HTTP handlers holding a *Service + writeJSON / writeError / writeServiceError helpers
    handler_test.go  # httptest handler tests + a fakeStore test double (individual tests, pre-refactor)
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
make test     # go test ./...
```

Or run `go test` directly, and add `-v` to see each named subtest:

```bash
go test ./... -v
```

> **Note:** This is the start branch. The handler tests already pass; they are just written as individual functions rather than as a table. Chapter 25 turns the repetitive cases into a single table-driven test.

## Contact

If you have any questions, feedback, or just want to connect, feel free to reach out to me on LinkedIn:

https://www.linkedin.com/in/mt26691/
