# Go Programming for Beginners: Build Real Backend Services

This repository contains the source code for the [Go Programming for Beginners: Build Real Backend Services](https://dalabs.academy) course.

## Prerequisites

- [Go](https://go.dev/dl/) 1.22 or newer (this repo was built with Go 1.26)
- [golangci-lint](https://golangci-lint.run/) (install with `brew install golangci-lint`)
- Comfort using a terminal and a code editor (VS Code with the Go extension recommended)
- [Git](https://git-scm.com/) installed

## Start Branch

```bash
git checkout 21-get-update-delete-task-start
```

The start branch is the finish state of Chapter 20. Two endpoints work: `POST /tasks` creates a task and returns `201`, and `GET /tasks` returns the list with `200`. `GET /tasks/{id}` is still a `501 Not Implemented` stub, and there is no way to update or delete a task. The store (`internal/task/store.go`) has only `Create` and `List`. This chapter finishes CRUD on a single resource: fill in the `get` handler and add update and delete.

## Finish Branch

```bash
git checkout 21-get-update-delete-task-finish
```

The finish branch completes CRUD on a single task:

- **`store.go`** — a sentinel `ErrNotFound` plus `Get(ctx, id) (Task, error)`, `Update(ctx, id, Task) (Task, error)`, and `Delete(ctx, id) error`, each guarded by the mutex and each returning `ErrNotFound` when the ID is absent. `Update` is a full replace that keeps the original `ID` and `CreatedAt`.
- **`handler.go`** — `get`, `update`, and `delete` parse the path ID with `strconv.Atoi` (bad ID → `400`), map `errors.Is(err, ErrNotFound)` → `404` and other errors → `500`, and register `GET`, `PUT`, and `DELETE` on `/tasks/{id}`. `DELETE` returns `204 No Content`.

## Lesson

[View the lesson on dalabs.academy](https://dalabs.academy/courses/go-programming-for-beginners-build-real-backend-services/building-the-task-api/get-update-delete-task)
<!-- After publishing, the /publish-chapter skill replaces the placeholder above with the actual URL -->

## Project Layout

```
cmd/
  server/
    main.go        # entry point: builds the mux, wires the store + handlers, starts the server
internal/
  task/
    task.go        # the Task resource (model)
    store.go       # in-memory Store (map[int]Task + sync.Mutex): Create, List
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

curl -i http://localhost:8080/tasks
# HTTP/1.1 200 OK

curl -i http://localhost:8080/tasks/1
# HTTP/1.1 501 Not Implemented   <- the stub this chapter replaces
```

```bash
make build   # go build ./... -> compiles every package
make fmt     # go fmt ./...   -> formats the code
make vet     # go vet ./...   -> reports suspicious code
make lint    # golangci-lint run -> runs the linter
make help    # list the available targets
```

> **Note:** This is the start branch. `GET /tasks/{id}` still returns `501`, and there is no update or delete yet. `go build ./...`, `go vet ./...`, and `golangci-lint run` are all clean.

## Contact

If you have any questions, feedback, or just want to connect, feel free to reach out to me on LinkedIn:

https://www.linkedin.com/in/mt26691/
