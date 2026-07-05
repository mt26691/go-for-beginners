# Go Programming for Beginners: Build Real Backend Services

This repository contains the source code for the [Go Programming for Beginners: Build Real Backend Services](https://dalabs.academy) course.

## Prerequisites

- [Go](https://go.dev/dl/) 1.22 or newer (this repo was built with Go 1.26)
- [golangci-lint](https://golangci-lint.run/) (install with `brew install golangci-lint`)
- Comfort using a terminal and a code editor (VS Code with the Go extension recommended)
- [Git](https://git-scm.com/) installed

## Start Branch

```bash
git checkout 22-validation-and-errors-start
```

The start branch is the finish state of Chapter 21: full CRUD works, but the API does not validate input and returns errors as bare plain text.

## Finish Branch

```bash
git checkout 22-validation-and-errors-finish
```

The finish branch never trusts input and makes every error look the same:

- **`task.go`** — a `MaxTitleLength` constant, two validation sentinels (`ErrTitleRequired`, `ErrTitleTooLong`), and a `Validate()` method on `*Task`. `Validate` trims whitespace from `Title` and `Description`, then rejects an empty title or one longer than `MaxTitleLength`. The checks are written by hand so the mechanics are visible; `github.com/go-playground/validator` is the library alternative, mentioned but not used.
- **`handler.go`** — a single `errorResponse` struct (`{"error":"..."}`) and a `writeError(w, status, message)` helper. Every bare `http.Error(...)` is gone: a malformed body, a bad id, and a failed validation all return `400` through the same JSON envelope. A `writeStoreError` helper maps `ErrNotFound` to `404` in one place and turns any other store error into a generic `500` (`"something went wrong"`) while logging the real cause server-side, so internal details never leak to the client.

Client mistakes (bad input) are `4xx`; only genuinely unexpected failures are `5xx`. This centralized error handling is what the recovery middleware in Chapter 32 builds on.

## Lesson

[View the lesson on dalabs.academy]({URL})
<!-- After publishing, the /publish-chapter skill replaces the placeholder above with the actual URL -->

## Project Layout

```
cmd/
  server/
    main.go        # entry point: builds the mux, wires the store + handlers, starts the server
internal/
  task/
    task.go        # the Task resource (model) + Validate()
    store.go       # in-memory Store (map[int]Task + sync.Mutex): Create, List, Get, Update, Delete
    handler.go     # task HTTP handlers + writeJSON / writeError / writeStoreError helpers
```

## Running the Server

```bash
make run     # go run ./cmd/server
```

Then in another terminal:

```bash
# an empty title is now rejected with a JSON error
curl -i -X POST http://localhost:8080/tasks -d '{"title":""}'
# HTTP/1.1 400 Bad Request
# Content-Type: application/json
# {"error":"title is required"}

# the happy path still works
curl -i -X POST http://localhost:8080/tasks -d '{"title":"Buy milk"}'
# HTTP/1.1 201 Created
# {"id":1,"title":"Buy milk","done":false,"createdAt":"..."}

# a malformed body is a 400, not a 500
curl -i -X POST http://localhost:8080/tasks -d 'not json'
# HTTP/1.1 400 Bad Request
# {"error":"invalid request body"}

# a missing task is now JSON too, not plain text
curl -i http://localhost:8080/tasks/999
# HTTP/1.1 404 Not Found
# Content-Type: application/json
# {"error":"task not found"}
```

```bash
make build   # go build ./... -> compiles every package
make vet     # go vet ./...   -> reports suspicious code
make lint    # golangci-lint run -> runs the linter
```

> **Note:** This is the finish branch. Input is validated, every error path returns the same JSON envelope with `Content-Type: application/json`, and `go build ./...`, `go vet ./...`, and `golangci-lint run` are all clean.

## Contact

If you have any questions, feedback, or just want to connect, feel free to reach out to me on LinkedIn:

https://www.linkedin.com/in/mt26691/
