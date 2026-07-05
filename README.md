# Go Programming for Beginners: Build Real Backend Services

This repository contains the source code for the [Go Programming for Beginners: Build Real Backend Services](https://dalabs.academy) course.

## Prerequisites

- [Go](https://go.dev/dl/) 1.22 or newer (this repo was built with Go 1.26)
- [golangci-lint](https://golangci-lint.run/) (install with `brew install golangci-lint`)
- Comfort using a terminal and a code editor (VS Code with the Go extension recommended)
- [Git](https://git-scm.com/) installed

## Start Branch

```bash
git checkout 20-create-and-list-tasks-start
```

The start branch is the compiling skeleton from Chapter 19. The project layout is in place: `cmd/server/main.go` boots a `http.NewServeMux()` and wires the store into the task handlers, `internal/task/task.go` holds the `Task` resource, `internal/task/store.go` holds an in-memory `Store` (a `map[int]Task` guarded by a `sync.Mutex`) with a `List` method, and `internal/task/handler.go` registers the routes. `GET /tasks` already returns an empty JSON array, but `POST /tasks` and `GET /tasks/{id}` are stubs that return `501 Not Implemented`. This chapter turns the `create` stub into a working handler and gives the store a `Create` method.

## Finish Branch

```bash
git checkout 20-create-and-list-tasks-finish
```

The finish branch has the first real endpoints working:

- **`store.go`** — the `Store` gains a `nextID` counter and a `Create(ctx, Task) (Task, error)` method that assigns the ID, stamps `CreatedAt` with `time.Now()`, and stores the task under the mutex. `List` now takes a `context.Context` and returns `([]Task, error)` so both methods share the shape the database-backed store will need in Section 6.
- **`handler.go`** — `create` decodes the JSON body with `json.NewDecoder`, calls `store.Create`, and returns `201 Created` with the stored task (including the server-assigned `id` and `createdAt`); a bad body returns `400`. `list` calls `store.List` and returns the array with `200`. Both handlers pass `r.Context()` down to the store. `get` is still a `501` stub — Chapter 21 fills it in.

## Lesson

[View the lesson on dalabs.academy](https://dalabs.academy/courses/go-programming-for-beginners-build-real-backend-services/building-the-task-api/create-and-list-tasks)
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
curl -i -X POST http://localhost:8080/tasks -d '{"title":"Buy milk"}'
# HTTP/1.1 201 Created
# Content-Type: application/json
# {"id":1,"title":"Buy milk","done":false,"createdAt":"2026-07-06T09:06:50.109193+10:00"}

curl -i http://localhost:8080/tasks
# HTTP/1.1 200 OK
# Content-Type: application/json
# [{"id":1,"title":"Buy milk","done":false,"createdAt":"2026-07-06T09:06:50.109193+10:00"}]
```

The server assigns the `id` and `createdAt`; the client only sends `title` and `description`. Input validation and a consistent JSON error shape arrive in Chapter 22, so a bad body returns a plain-text `400` for now.

```bash
make build   # go build ./... -> compiles every package
make fmt     # go fmt ./...   -> formats the code
make vet     # go vet ./...   -> reports suspicious code
make lint    # golangci-lint run -> runs the linter
make help    # list the available targets
```

> **Note:** This is the finish branch. `POST /tasks` creates a task and returns `201`, `GET /tasks` returns the list with `200`, and `GET /tasks/{id}` is still a `501` stub (Chapter 21). `go build ./...`, `go vet ./...`, and `golangci-lint run` are all clean.

## Contact

If you have any questions, feedback, or just want to connect, feel free to reach out to me on LinkedIn:

https://www.linkedin.com/in/mt26691/
