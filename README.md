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

The start branch is the finish state of Chapter 24: `internal/task/handler_test.go` tests the HTTP handlers with `httptest` and a `fakeStore`, written as **individual functions and `t.Run` subtests** that repeat the same request/record/assert shape.

## Finish Branch

```bash
git checkout 25-table-driven-tests-finish
```

The finish branch refactors those repetitive tests into the idiomatic **table-driven** form. Only `internal/task/handler_test.go` changes:

- **A table of cases.** `TestCreateTask` becomes a `[]struct{ name; title; wantStatus; wantError string }` iterated with `t.Run(tc.name, ...)`, so each row is one named subtest with pinpoint failure output. `TestGetTask` gets its own small table (found, non-numeric id, missing id) rather than being forced into the create table.
- **More cases, fewer functions.** The table adds a **too-long title → 400** case and a **non-numeric id → 400** case that the individual tests never covered, while collapsing the repeated boilerplate.
- **Good failure messages.** Every assertion reports got vs want (`t.Errorf("status = %d, want %d", ...)`) and checks the `{"error": ...}` envelope where relevant.
- **`t.Parallel()`.** Each subtest runs in parallel on its own fresh `fakeStore`, so state never leaks. Since Go 1.22 each loop iteration gets a fresh `tc`, so the old `tc := tc` copy is no longer needed.
- **`TestListTasks` stays a single straight-line test** — one scenario, not a family of cases, so a table would only add ceremony.

The result is denser test code that covers more of the handler surface. Run `go test ./... -cover` to see the coverage figure.

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
    handler_test.go  # table-driven httptest handler tests + a fakeStore test double (Chapter 25)
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

Or run `go test` directly, and add `-v` to see each named subtest run (in parallel), or `-cover` for the coverage figure:

```bash
go test ./... -v
go test ./... -cover
```

> **Note:** This is the finish branch. `internal/task/handler_test.go` is now table-driven and covers more cases (too-long title, non-numeric id) than Chapter 24 did. `go test ./...` passes, and `go build ./...`, `go vet ./...`, and `golangci-lint run` are all clean.

## Contact

If you have any questions, feedback, or just want to connect, feel free to reach out to me on LinkedIn:

https://www.linkedin.com/in/mt26691/
