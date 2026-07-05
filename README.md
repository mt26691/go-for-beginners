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

The start branch is the finish state of Chapter 20. `POST /tasks` creates a task and returns `201`, `GET /tasks` returns the list with `200`, and `GET /tasks/{id}` is still a `501 Not Implemented` stub. The store (`internal/task/store.go`) has `Create` and `List` methods; it has no way to fetch, update, or delete a single task yet. This chapter fills in the `get` handler and adds update and delete to finish CRUD on a single resource.

## Finish Branch

```bash
git checkout 21-get-update-delete-task-finish
```

The finish branch completes CRUD on a single task:

- **`store.go`** — a sentinel `ErrNotFound` (`var ErrNotFound = errors.New("task not found")`) plus three new mutex-guarded methods: `Get(ctx, id) (Task, error)`, `Update(ctx, id, Task) (Task, error)`, and `Delete(ctx, id) error`. Each returns `ErrNotFound` when the ID is absent. `Update` is a **full replace**: it keeps the original `ID` and `CreatedAt` and takes every other field from the incoming task.
- **`handler.go`** — `get`, `update`, and `delete` read the path with `r.PathValue("id")` and parse it with `strconv.Atoi`; a non-numeric ID returns `400`. Handlers map `errors.Is(err, ErrNotFound)` to `404` and any other error to `500`. `get` returns `200` with the task, `update` returns `200` with the updated task, and `delete` returns `204 No Content` with an empty body. `Routes` now registers `GET`, `PUT`, and `DELETE` on `/tasks/{id}`.

`PUT` is a full replace. The partial-update alternative is `PATCH`, which this course does not implement. Validation and a consistent JSON error shape still arrive in Chapter 22, so a bad ID or bad body returns a plain-text `400` for now.

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
    store.go       # in-memory Store (map[int]Task + sync.Mutex): Create, List, Get, Update, Delete
    handler.go     # task HTTP handlers + writeJSON helper
```

`internal/` is a Go convention: packages under it can only be imported by code inside this module, so the task domain stays private to this service. Keep the package count small — this is a small API, not a place for premature "clean architecture."

## Running the Server

```bash
make run     # go run ./cmd/server
```

Then in another terminal:

```bash
# create a task so there is an id 1 to work with
curl -i -X POST http://localhost:8080/tasks -d '{"title":"Write the chapter","description":"Draft chapter 21"}'
# HTTP/1.1 201 Created

curl -i http://localhost:8080/tasks/1
# HTTP/1.1 200 OK
# {"id":1,"title":"Write the chapter","description":"Draft chapter 21","done":false,"createdAt":"..."}

curl -i http://localhost:8080/tasks/999
# HTTP/1.1 404 Not Found
# task not found

curl -i -X PUT http://localhost:8080/tasks/1 -d '{"title":"Ship the chapter","done":true}'
# HTTP/1.1 200 OK  (full replace: id and createdAt are preserved)
# {"id":1,"title":"Ship the chapter","done":true,"createdAt":"..."}

curl -i -X DELETE http://localhost:8080/tasks/1
# HTTP/1.1 204 No Content  (empty body)

curl -i http://localhost:8080/tasks/abc
# HTTP/1.1 400 Bad Request
# invalid task id
```

```bash
make build   # go build ./... -> compiles every package
make fmt     # go fmt ./...   -> formats the code
make vet     # go vet ./...   -> reports suspicious code
make lint    # golangci-lint run -> runs the linter
make help    # list the available targets
```

> **Note:** This is the finish branch. All five task routes work — `POST /tasks`, `GET /tasks`, `GET /tasks/{id}`, `PUT /tasks/{id}`, and `DELETE /tasks/{id}`. `go build ./...`, `go vet ./...`, and `golangci-lint run` are all clean.

## Contact

If you have any questions, feedback, or just want to connect, feel free to reach out to me on LinkedIn:

https://www.linkedin.com/in/mt26691/
