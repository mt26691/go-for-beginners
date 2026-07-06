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

The finish branch adds the first tests. A new file, `internal/task/handler_test.go`, exercises the HTTP handlers in-process with `net/http/httptest`:

- **A `fakeStore` test double** implements the `TaskStore` interface with a plain map. It is injected via `task.NewService(fakeStore)` → `task.NewHandler(svc)`, so the tests run entirely in memory with no database — the payoff of the Chapter 23 interface seam.
- **`httptest.NewRequest` + `httptest.NewRecorder`** drive each handler through the real mux and assert on `rec.Code`, the decoded JSON body, and the `Content-Type` header.
- **Cases covered:** `POST /tasks` valid → 201 + task, `POST /tasks` empty title → 400 + `{"error":"title is required"}`, `GET /tasks` → 200 array, `GET /tasks/{id}` found → 200, `GET /tasks/999` → 404. Each test builds a fresh `fakeStore`, so state never leaks between cases.
- **A `make test` target** runs `go test ./...`.

The tests are written as individual functions with `t.Run` subtests; Chapter 25 refactors the repetitive cases into a single table-driven test.

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
    task.go          # the Task resource (model) + Validate()
    store.go         # in-memory Store (map[int]Task + sync.Mutex) + compile-time TaskStore assertion
    service.go       # TaskStore interface + Service (business logic) between handlers and storage
    handler.go       # task HTTP handlers holding a *Service + writeJSON / writeError / writeServiceError helpers
    handler_test.go  # httptest handler tests + a fakeStore test double (Chapter 24)
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

> **Note:** This is the finish branch. `internal/task/handler_test.go` tests the HTTP handlers in-process with `httptest` and a fake store, so the suite runs in milliseconds without a database. `go test ./...` passes, and `go build ./...`, `go vet ./...`, and `golangci-lint run` are all clean.

## Contact

If you have any questions, feedback, or just want to connect, feel free to reach out to me on LinkedIn:

https://www.linkedin.com/in/mt26691/
