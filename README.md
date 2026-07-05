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

The start branch is the empty target layout. The single-file `main.go` from Section 3 is gone; in its place is an idiomatic Go project layout: `cmd/server/main.go` holds a minimal entry point that boots a `http.NewServeMux()` with just `GET /ping`, and `internal/task/task.go` holds the finalized `Task` resource (`id`, `title`, `description`, `done`, `createdAt`). The chapter fills in the `internal/task` package — an in-memory `Store` and the task HTTP handlers — and wires them into `main`.

## Finish Branch

```bash
git checkout 19-designing-the-api-finish
```

The finish branch is a compiling skeleton. `internal/task` grows two files:

- **`store.go`** — a `Store` wrapping a `map[int]Task` guarded by a `sync.Mutex` (every request runs in its own goroutine), with a `NewStore` constructor and a `List` method. No `nextID` yet and no `Create`/`Get` methods — those arrive in Chapters 20 and 21.
- **`handler.go`** — a `Handler` holding the `*Store`, a `Routes(mux)` method that registers `GET /tasks`, `POST /tasks`, and `GET /tasks/{id}`, and the shared `writeJSON` helper. `list` returns whatever the store holds (an empty `[]` for now); `create` and `get` are deliberate stubs that return `501 Not Implemented`.

`cmd/server/main.go` wires it together: build the mux, register `GET /ping`, create the store, hand it to `NewHandler`, register the task routes, and start listening.

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
    store.go       # in-memory Store (map[int]Task + sync.Mutex)
    handler.go     # task HTTP handlers + writeJSON helper
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

curl -i http://localhost:8080/tasks
# HTTP/1.1 200 OK
# Content-Type: application/json
# []

curl -i -X POST http://localhost:8080/tasks -d '{"title":"Buy milk"}'
# HTTP/1.1 501 Not Implemented
# not implemented

curl -i http://localhost:8080/tasks/1
# HTTP/1.1 501 Not Implemented
# not implemented
```

`GET /tasks` returns an empty JSON array `[]` (not `null`) from the store; the two stubbed routes return `501` until Chapters 20 and 21 fill them in.

```bash
make build   # go build ./... -> compiles every package
make fmt     # go fmt ./...   -> formats the code
make vet     # go vet ./...   -> reports suspicious code
make lint    # golangci-lint run -> runs the linter
make help    # list the available targets
```

> **Note:** This is the finish branch — a compiling skeleton. The layout is in place, `GET /tasks` returns `[]` from the in-memory store, and `POST /tasks` and `GET /tasks/{id}` return `501 Not Implemented`. `go build ./...`, `go vet ./...`, and `golangci-lint run` are all clean.

## Contact

If you have any questions, feedback, or just want to connect, feel free to reach out to me on LinkedIn:

https://www.linkedin.com/in/mt26691/
