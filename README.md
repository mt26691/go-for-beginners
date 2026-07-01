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

The finish branch adds three focused context demos on top of the JSON server:

- **Cancellation** — `slowHandler` (`GET /slow`) starts a 5-second job and `select`s on `r.Context().Done()`. When the client disconnects, the request's context is cancelled automatically and the handler logs `context canceled` instead of finishing.
- **Deadlines** — `getTask` (`GET /tasks/{id}`) derives `ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)`, always `defer cancel()`s, and passes `ctx` into a `findTask(ctx, id)` store stand-in. A dedicated `GET /slow-store` route uses a 100ms deadline against the same ~500ms "query" so it trips `context.DeadlineExceeded` and returns `504`.
- **Request-scoped values** — `traceHandler` (`GET /trace`) reads an `X-Request-ID` header, stashes it on the context with a typed key, and reads it back downstream through a `requestIDFrom(ctx)` accessor.

`findTask` is a stand-in for the real `pgx` queries added in Chapter 27; it takes a `context.Context` as its first argument for exactly the same reason the database calls will.

## Lesson

[View the lesson on dalabs.academy](https://dalabs.academy/courses/go-programming-for-beginners-build-real-backend-services/your-first-http-server/context)
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
# {"id":42,"title":"Sample task","done":false}   (returns after ~500ms; the 2s deadline is not hit)

# 1. Client-disconnect cancellation: kill the request early and watch the server log
curl --max-time 1 http://localhost:8080/slow
# curl: (28) Operation timed out after 1005 milliseconds with 0 bytes received
# server log: slow: request cancelled: context canceled

# 2. Caller-imposed deadline: the 100ms deadline fires before the 500ms "query"
curl -i http://localhost:8080/slow-store
# HTTP/1.1 504 Gateway Timeout
# context deadline exceeded
# server log: slow-store: context deadline exceeded

# 3. Request-scoped value carried on the context
curl -H "X-Request-ID: abc-123" http://localhost:8080/trace
# {"requestID":"abc-123"}
# server log: trace: handling request abc-123
```

```bash
make build   # go build   -> compiles the binary
make fmt     # go fmt ./... -> formats the code
make vet     # go vet ./... -> reports suspicious code
make lint    # golangci-lint run -> runs the linter
make help    # list the available targets
```

> **Note:** This is the finish branch — `main.go` keeps the Chapter 17 JSON server (`Task`, `writeJSON`, and the task routes) and adds the `context.Context` demos: `slowHandler`, `findTask`, `slowStoreHandler`, and `traceHandler`. `go build`, `go vet`, and `gofmt` are all clean. `make lint` still reports one `errcheck` finding on the unchecked `fmt.Fprintln` in `ping` (unchanged since Chapter 14) — that is expected for this build-and-run branch.

## Contact

If you have any questions, feedback, or just want to connect, feel free to reach out to me on LinkedIn:

https://www.linkedin.com/in/mt26691/
