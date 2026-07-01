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

The start branch carries over the Chapter 16 server. The routes are wired on an explicit `http.NewServeMux()` with Go 1.22 method-plus-path patterns, but the task handlers still fake their JSON: `listTasks` writes a hard-coded JSON *string*, `createTask` returns `201` with plain text, and `getTask` echoes the id as plain text.

## Finish Branch

```bash
git checkout 17-working-with-json-finish
```

The finish branch turns those fakes into real JSON with `encoding/json`. It adds a `Task` struct with JSON struct tags (`json:"title"`, `omitempty`, and one field excluded with `json:"-"`) and a reusable `writeJSON(w, status, v)` helper that sets `Content-Type: application/json`, writes the status, and streams the value with `json.NewEncoder(w).Encode`. `listTasks` encodes a real `[]Task`, `getTask` parses the `{id}` and returns a single `Task`, and `createTask` decodes the request body with `json.NewDecoder(r.Body).Decode` and echoes it back with `201` (or `400` on a decode error). `GET /ping` stays plain text for contrast.

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
# [{"id":1,"title":"Write the JSON chapter","done":true},{"id":2,"title":"Record the demo","description":"curl every route","done":false}]

curl http://localhost:8080/tasks/42
# {"id":42,"title":"Sample task","done":false}

curl -X POST http://localhost:8080/tasks -d '{"title":"Buy milk","done":false}'
# {"id":0,"title":"Buy milk","done":false}

curl -i -X POST http://localhost:8080/tasks -d 'not json'
# HTTP/1.1 400 Bad Request
# invalid JSON body
```

```bash
make build   # go build   -> compiles the binary
make fmt     # go fmt ./... -> formats the code
make vet     # go vet ./... -> reports suspicious code
make lint    # golangci-lint run -> runs the linter
make help    # list the available targets
```

> **Note:** This is the finish branch — `main.go` contains the complete Chapter 17 server. Every task route now speaks real JSON through the `writeJSON` helper and `json.NewDecoder`. Notice how `description` is omitted when empty (`omitempty`) and the `Internal` field never appears on the wire (`json:"-"`). The `Task` struct and `writeJSON` helper here are reused by the rest of the course.

## Contact

If you have any questions, feedback, or just want to connect, feel free to reach out to me on LinkedIn:

https://www.linkedin.com/in/mt26691/
