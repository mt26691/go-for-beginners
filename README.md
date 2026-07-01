# Go Programming for Beginners: Build Real Backend Services

This repository contains the source code for the [Go Programming for Beginners: Build Real Backend Services](https://dalabs.academy) course.

## Prerequisites

- [Go](https://go.dev/dl/) 1.22 or newer (this repo was built with Go 1.26)
- [golangci-lint](https://golangci-lint.run/) (install with `brew install golangci-lint`)
- Comfort using a terminal and a code editor (VS Code with the Go extension recommended)
- [Git](https://git-scm.com/) installed

## Start Branch

```bash
git checkout 17-working-with-json-start
```

The start branch carries over the Chapter 16 server. The routes are wired on an explicit `http.NewServeMux()` with Go 1.22 method-plus-path patterns (`GET /ping`, `GET /tasks`, `POST /tasks`, `GET /tasks/{id}`), but the task handlers still fake their JSON: `listTasks` writes a hard-coded JSON *string*, `createTask` returns `201` with plain text, and `getTask` echoes the id as plain text. Your goal is to make these handlers speak real JSON with `encoding/json`.

## Finish Branch

```bash
git checkout 17-working-with-json-finish
```

The finish branch introduces a `Task` struct with JSON struct tags, a reusable `writeJSON` helper, and real encode/decode: `listTasks` encodes a `[]Task`, `getTask` returns a single `Task` built from the `{id}`, and `createTask` decodes the request body into a `Task` and echoes it back with `201` (or `400` on a decode error).

## Lesson

[View the lesson on dalabs.academy](https://dalabs.academy/courses/go-programming-for-beginners-build-real-backend-services/your-first-http-server/working-with-json)
<!-- After publishing, the /publish-chapter skill replaces the placeholder above with the actual URL -->

## Running the Server

```bash
go run .
```

Then in another terminal try each route:

```bash
curl http://localhost:8080/ping
# pong

curl http://localhost:8080/tasks
# [{"id":1,"title":"Write the routing chapter"}]

curl -i -X POST http://localhost:8080/tasks
# HTTP/1.1 201 Created
# task created

curl http://localhost:8080/tasks/42
# task 42
```

```bash
make build   # go build   -> compiles the binary
make fmt     # go fmt ./... -> formats the code
make vet     # go vet ./... -> reports suspicious code
make lint    # golangci-lint run -> runs the linter
make help    # list the available targets
```

> **Note:** This is the start branch — the task handlers still fake their responses (a hard-coded JSON string and plain text). Over the chapter you will replace them with real JSON encoding and decoding using the `encoding/json` package.

## Contact

If you have any questions, feedback, or just want to connect, feel free to reach out to me on LinkedIn:

https://www.linkedin.com/in/mt26691/
