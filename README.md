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

The start branch is the finish state of Chapter 21: full CRUD on a single task works, but the API trusts its input and returns errors as bare plain text.

- `POST /tasks` stores whatever body it receives — even `{"title":""}` — because nothing validates the decoded task.
- Every error path calls `http.Error(w, "...", status)`, which writes a `text/plain` body. So a bad id, a malformed body, and a missing task all return prose, not JSON, and a JSON client has to special-case them.

This chapter validates the incoming task (a title is required and capped in length) and replaces every bare `http.Error` with a single JSON error envelope so all errors share one shape.

## Finish Branch

```bash
git checkout 22-validation-and-errors-finish
```

The finish branch validates input and returns a consistent JSON error envelope everywhere.

## Lesson

[View the lesson on dalabs.academy](<!-- dalabs:22-validation-and-errors -->)
<!-- After publishing, the /publish-chapter skill replaces the placeholder above with the actual URL -->

## Project Layout

```
cmd/
  server/
    main.go        # entry point: builds the mux, wires the store + handlers, starts the server
internal/
  task/
    task.go        # the Task resource (model)
    store.go       # in-memory Store (map[int]Task + sync.Mutex): Create, List, Get, Update, Delete
    handler.go     # task HTTP handlers + writeJSON helper
```

## Running the Server

```bash
make run     # go run ./cmd/server
```

Then in another terminal:

```bash
# the input is never validated: an empty title is accepted and returns 201
curl -i -X POST http://localhost:8080/tasks -d '{"title":""}'
# HTTP/1.1 201 Created

# errors come back as plain text, not JSON
curl -i http://localhost:8080/tasks/999
# HTTP/1.1 404 Not Found
# Content-Type: text/plain; charset=utf-8
# task not found
```

```bash
make build   # go build ./... -> compiles every package
make vet     # go vet ./...   -> reports suspicious code
make lint    # golangci-lint run -> runs the linter
```

> **Note:** This is the start branch (the finish state of Chapter 21). It builds, vets, and lints clean, but it does no input validation and returns plain-text errors. Chapter 22 adds validation and a JSON error envelope.

## Contact

If you have any questions, feedback, or just want to connect, feel free to reach out to me on LinkedIn:

https://www.linkedin.com/in/mt26691/
