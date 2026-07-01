# Go Programming for Beginners: Build Real Backend Services

This repository contains the source code for the [Go Programming for Beginners: Build Real Backend Services](https://dalabs.academy) course.

## Prerequisites

- [Go](https://go.dev/dl/) 1.22 or newer (this repo was built with Go 1.26)
- [golangci-lint](https://golangci-lint.run/) (install with `brew install golangci-lint`)
- Comfort using a terminal and a code editor (VS Code with the Go extension recommended)
- [Git](https://git-scm.com/) installed

## Start Branch

```bash
git checkout 18-context-start
```

The start branch carries over the Chapter 17 JSON server: a `Task` struct with JSON struct tags, the reusable `writeJSON` helper, and the `GET /ping`, `GET /tasks`, `POST /tasks`, `GET /tasks/{id}` routes on an explicit `http.NewServeMux()`. Your goal in this chapter is to add a spotlight on `context.Context` — carrying a cancellation signal, a deadline, and request-scoped values from the handler down the call chain.

## Finish Branch

```bash
git checkout 18-context-finish
```

The finish branch adds three focused context demos on top of the JSON server: a `GET /slow` handler that watches `r.Context().Done()` so a disconnected client cancels the request, a `findTask(ctx, id)` store stand-in that `getTask` calls through a `context.WithTimeout` deadline (with a `GET /slow-store` route that trips `context.DeadlineExceeded`), and a `GET /trace` handler that stashes a request ID on the context with a typed key and reads it back downstream.

## Lesson

[View the lesson on dalabs.academy](<!-- dalabs:18-context -->)
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
```

```bash
make build   # go build   -> compiles the binary
make fmt     # go fmt ./... -> formats the code
make vet     # go vet ./... -> reports suspicious code
make lint    # golangci-lint run -> runs the linter
make help    # list the available targets
```

> **Note:** This is the start branch — `main.go` is exactly the Chapter 17 JSON server, with no context handling yet. Follow the chapter to add the `/slow`, `/slow-store`, and `/trace` demos and to wire a `context.WithTimeout` deadline into `getTask`.

## Contact

If you have any questions, feedback, or just want to connect, feel free to reach out to me on LinkedIn:

https://www.linkedin.com/in/mt26691/
