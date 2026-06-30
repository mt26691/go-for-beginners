# Go Programming for Beginners: Build Real Backend Services

This repository contains the source code for the [Go Programming for Beginners: Build Real Backend Services](https://dalabs.academy) course.

## Prerequisites

- [Go](https://go.dev/dl/) 1.22 or newer (this repo was built with Go 1.26)
- [golangci-lint](https://golangci-lint.run/) (install with `brew install golangci-lint`)
- Comfort using a terminal and a code editor (VS Code with the Go extension recommended)
- [Git](https://git-scm.com/) installed

## Start Branch

```bash
git checkout 16-routing-with-servemux-start
```

The start branch carries over the Chapter 15 server, which registers `helloHandler`, `greetHandler`, `echoHandler`, and `pingHandler` on the **default mux** with bare path strings (`http.HandleFunc("/greet", ...)`). Your goal is to move to an explicit `http.NewServeMux()` and route the Task API surface with Go 1.22 method-plus-path patterns.

## Finish Branch

```bash
git checkout 16-routing-with-servemux-finish
```

The finish branch creates an explicit `http.NewServeMux()` and registers method-plus-path patterns: `GET /ping`, `GET /tasks`, `POST /tasks`, and `GET /tasks/{id}`. The wildcard handler reads the id with `r.PathValue("id")`. The task handlers are stubs for now (the in-memory store arrives in Section 4), but every route returns a real, curl-able response. `http.ListenAndServe` is passed `mux` instead of `nil`, so the mux you built owns the routing.

## Lesson

[View the lesson on dalabs.academy](https://dalabs.academy/courses/go-programming-for-beginners-build-real-backend-services/your-first-http-server/routing-with-servemux)
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

curl -i -X DELETE http://localhost:8080/tasks
# HTTP/1.1 405 Method Not Allowed
# Allow: GET, HEAD, POST

curl -i http://localhost:8080/nope
# HTTP/1.1 404 Not Found
```

```bash
make build   # go build   -> compiles the binary
make fmt     # go fmt ./... -> formats the code
make vet     # go vet ./... -> reports suspicious code
make lint    # golangci-lint run -> runs the linter
make help    # list the available targets
```

> **Note:** This is the finish branch — `main.go` contains the complete Chapter 16 server using an explicit `http.NewServeMux()` with Go 1.22 method-plus-path patterns. Run `go run .` and curl each route to see the routing, the automatic `405` on a method mismatch, and the `404` on an unknown path.

## Contact

If you have any questions, feedback, or just want to connect, feel free to reach out to me on LinkedIn:

https://www.linkedin.com/in/mt26691/
