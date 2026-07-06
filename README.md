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

The start branch is the finish state of Chapter 22: full CRUD works and every error returns the same JSON envelope, but the HTTP handlers talk to the concrete `*Store` directly. There is no service layer and no storage interface.

## Finish Branch

```bash
git checkout 23-organizing-the-code-finish
```

The finish branch splits the code into three layers behind an interface — and every HTTP response stays byte-for-byte identical to Chapter 22:

- **`handler.go`** — the HTTP layer. `Handler` now holds a `*Service` instead of a `*Store`, and `NewHandler(svc *Service)` takes it as a constructor argument (dependency injection). The handlers only decode requests, call the service, and encode responses; they no longer call `Validate()` or the store themselves.
- **`service.go`** — the new business-logic layer. It defines the `TaskStore` interface (the storage seam foreshadowed in Chapter 11) and a `Service` that holds a `TaskStore` — the interface, not the concrete store. `Service.Create`/`Update` run `Validate()`, then delegate to the store. `NewService(store TaskStore)` injects the store.
- **`store.go`** — the in-memory `Store` is unchanged apart from a compile-time assertion, `var _ TaskStore = (*Store)(nil)`, that proves it satisfies the interface. If a method ever drifts, the build breaks here instead of at the call site.
- **`main.go`** — wires the layers with three lines: `store := task.NewStore(); svc := task.NewService(store); h := task.NewHandler(svc)`.

The point is the seam: handlers depend on the interface via the service, not on the concrete store. It pays off twice — Section 5 tests the handlers against a fake store, and Section 6 swaps in a PostgreSQL store with a tiny diff. Nothing in `handler.go` or `service.go` changes when the storage does.

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
    task.go        # the Task resource (model) + Validate()
    store.go       # in-memory Store (map[int]Task + sync.Mutex) + compile-time TaskStore assertion
    service.go     # TaskStore interface + Service (business logic) between handlers and storage
    handler.go     # task HTTP handlers holding a *Service + writeJSON / writeError / writeServiceError helpers
```

## Running the Server

```bash
make run     # go run ./cmd/server
```

Then in another terminal — the responses are the same as Chapter 22, because this is a pure refactor:

```bash
# an empty title is rejected with a JSON error
curl -i -X POST http://localhost:8080/tasks -d '{"title":""}'
# HTTP/1.1 400 Bad Request
# Content-Type: application/json
# {"error":"title is required"}

# the happy path returns 201 with the created task
curl -i -X POST http://localhost:8080/tasks -d '{"title":"Buy milk"}'
# HTTP/1.1 201 Created
# {"id":1,"title":"Buy milk","done":false,"createdAt":"..."}

# list returns a JSON array
curl -i http://localhost:8080/tasks
# HTTP/1.1 200 OK
# [{"id":1,"title":"Buy milk","done":false,"createdAt":"..."}]

# a missing task is a JSON 404
curl -i http://localhost:8080/tasks/999
# HTTP/1.1 404 Not Found
# Content-Type: application/json
# {"error":"task not found"}

# delete returns 204 No Content
curl -i -X DELETE http://localhost:8080/tasks/1
# HTTP/1.1 204 No Content
```

```bash
make build   # go build ./... -> compiles every package
make vet     # go vet ./...   -> reports suspicious code
make lint    # golangci-lint run -> runs the linter
```

> **Note:** This is the finish branch. The code is now layered handler -> service -> `TaskStore` interface <- in-memory `Store`, wired by dependency injection, and every HTTP response is identical to Chapter 22. `go build ./...`, `go vet ./...`, and `golangci-lint run` are all clean.

## Contact

If you have any questions, feedback, or just want to connect, feel free to reach out to me on LinkedIn:

https://www.linkedin.com/in/mt26691/
