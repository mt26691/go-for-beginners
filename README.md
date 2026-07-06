# Go Programming for Beginners: Build Real Backend Services

This repository contains the source code for the [Go Programming for Beginners: Build Real Backend Services](https://dalabs.academy) course.

## Prerequisites

- [Go](https://go.dev/dl/) 1.22 or newer (this repo was built with Go 1.26)
- [golangci-lint](https://golangci-lint.run/) (install with `brew install golangci-lint`)
- Comfort using a terminal and a code editor (VS Code with the Go extension recommended)
- [Git](https://git-scm.com/) installed

## Start Branch

```bash
git checkout 23-organizing-the-code-start
```

The start branch is the finish state of Chapter 22: full CRUD works and every error returns the same JSON envelope, but everything lives in one flat `internal/task` package and the HTTP handlers talk to the concrete `*Store` directly. There is no service layer and no storage interface yet.

## Finish Branch

```bash
git checkout 23-organizing-the-code-finish
```

The finish branch splits the code into three layers behind an interface, with the HTTP behavior left completely unchanged.

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
    handler.go     # task HTTP handlers holding a *Store directly + writeJSON / writeError / writeStoreError helpers
```

## Running the Server

```bash
make run     # go run ./cmd/server
```

Then in another terminal:

```bash
curl -i -X POST http://localhost:8080/tasks -d '{"title":""}'         # 400 {"error":"title is required"}
curl -i -X POST http://localhost:8080/tasks -d '{"title":"Buy milk"}'  # 201 Created
curl -i http://localhost:8080/tasks                                    # 200 [ ... ]
curl -i http://localhost:8080/tasks/999                                # 404 {"error":"task not found"}
```

```bash
make build   # go build ./... -> compiles every package
make vet     # go vet ./...   -> reports suspicious code
make lint    # golangci-lint run -> runs the linter
```

> **Note:** This is the start branch. The API already works end to end, but the handlers depend on the concrete `*Store`. In this chapter we introduce a service layer and a `TaskStore` interface so the storage can change later — without changing a single HTTP response.

## Contact

If you have any questions, feedback, or just want to connect, feel free to reach out to me on LinkedIn:

https://www.linkedin.com/in/mt26691/
