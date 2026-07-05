# Go Programming for Beginners: Build Real Backend Services

This repository contains the source code for the [Go Programming for Beginners: Build Real Backend Services](https://dalabs.academy) course.

## Prerequisites

- [Go](https://go.dev/dl/) 1.22 or newer (this repo was built with Go 1.26)
- [golangci-lint](https://golangci-lint.run/) (install with `brew install golangci-lint`)
- Comfort using a terminal and a code editor (VS Code with the Go extension recommended)
- [Git](https://git-scm.com/) installed

## Start Branch

```bash
git checkout 19-designing-the-api-start
```

The start branch is the empty target layout. The single-file `main.go` from Section 3 is gone; in its place is an idiomatic Go project layout: `cmd/server/main.go` holds a minimal entry point that boots a `http.NewServeMux()` with just `GET /ping`, and `internal/task/task.go` holds the finalized `Task` resource (`id`, `title`, `description`, `done`, `createdAt`). Your goal in this chapter is to fill in the `internal/task` package — an in-memory `Store` and the task HTTP handlers — and wire them into `main`.

## Finish Branch

```bash
git checkout 19-designing-the-api-finish
```

The finish branch is a compiling skeleton: `internal/task` grows a `Store` (a `map[int]Task` guarded by a `sync.Mutex`) with a `NewStore` constructor and a `List` method, plus a `Handler` that registers `GET /tasks`, `POST /tasks`, and `GET /tasks/{id}` on the mux. `main` wires the store into the handler. `GET /tasks` returns `[]` from the store; `POST /tasks` and `GET /tasks/{id}` are deliberate stubs that return `501 Not Implemented` — the real create/read logic lands in Chapters 20 and 21.

## Lesson

[View the lesson on dalabs.academy](https://dalabs.academy/courses/go-programming-for-beginners-build-real-backend-services/building-the-task-api/designing-the-api)
<!-- After publishing, the /publish-chapter skill replaces the placeholder above with the actual URL -->

## Project Layout

```
cmd/
  server/
    main.go        # entry point: builds the mux, wires the store + handlers, starts the server
internal/
  task/
    task.go        # the Task resource (model)
    store.go       # in-memory Store (finish branch)
    handler.go     # task HTTP handlers (finish branch)
```

`internal/` is a Go convention: packages under it can only be imported by code inside this module, so the task domain stays private to this service. Keep the package count small — this is a small API, not a place for premature "clean architecture."

## Running the Server

```bash
make run     # go run ./cmd/server
```

Then in another terminal:

```bash
curl http://localhost:8080/ping
# pong
```

On the start branch there are no task routes yet, so `curl http://localhost:8080/tasks` returns `404`.

```bash
make build   # go build ./... -> compiles every package
make fmt     # go fmt ./...   -> formats the code
make vet     # go vet ./...   -> reports suspicious code
make lint    # golangci-lint run -> runs the linter
make help    # list the available targets
```

> **Note:** This is the start branch — the layout scaffold only. `cmd/server/main.go` boots a `/ping`-only server and `internal/task/task.go` defines the `Task` struct. `go build ./...`, `go vet ./...`, and `golangci-lint run` are all clean. Follow the chapter to add `internal/task/store.go` and `internal/task/handler.go` and wire the task routes.

## Contact

If you have any questions, feedback, or just want to connect, feel free to reach out to me on LinkedIn:

https://www.linkedin.com/in/mt26691/
