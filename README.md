# Go Programming for Beginners: Build Real Backend Services

This repository contains the source code for the [Go Programming for Beginners: Build Real Backend Services](https://dalabs.academy) course.

## Prerequisites

- [Go](https://go.dev/dl/) 1.22 or newer (this repo was built with Go 1.26)
- [golangci-lint](https://golangci-lint.run/) (install with `brew install golangci-lint`)
- Comfort using a terminal and a code editor (VS Code with the Go extension recommended)
- [Git](https://git-scm.com/) installed

## Start Branch

```bash
git checkout 15-handlers-and-responses-start
```

The start branch carries over the working "Hello from Go!" server from Chapter 14. Your goal is to add `greetHandler`, `echoHandler`, and `pingHandler`, then register them on the mux to reach the finish state.

## Finish Branch

```bash
git checkout 15-handlers-and-responses-finish
```

The finish branch extends the Chapter 14 server with four handlers that demonstrate the full request/response API. `greetHandler` reads a `name` query parameter and returns 400 when it is missing. `echoHandler` reads the request body with `io.ReadAll`, sets `Content-Type`, and returns 201. `pingHandler` is a custom struct satisfying `http.Handler` via `ServeHTTP`. All four are registered on the default mux.

## Lesson

[View the lesson on dalabs.academy](<!-- dalabs:15-handlers-and-responses -->)
<!-- After publishing, the /publish-chapter skill replaces the placeholder above with the actual URL -->

## Running the Server

```bash
go run .
```

Then in another terminal try each route:

```bash
curl http://localhost:8080/
# Hello from Go!

curl 'http://localhost:8080/greet?name=Tung'
# Hello, Tung!

curl -i 'http://localhost:8080/greet'
# HTTP/1.1 400 Bad Request
# missing 'name' query parameter

curl -i -X POST --data 'hello body' http://localhost:8080/echo
# HTTP/1.1 201 Created
# hello body

curl http://localhost:8080/ping
# pong
```

```bash
make build   # go build   -> compiles the binary
make fmt     # go fmt ./... -> formats the code
make vet     # go vet ./... -> reports suspicious code
make lint    # golangci-lint run -> runs the linter
make help    # list the available targets
```

> **Note:** This is the finish branch — `main.go` contains the complete Chapter 15 server with all four handlers. Run `go run .` and curl each route to verify the responses.

## Contact

If you have any questions, feedback, or just want to connect, feel free to reach out to me on LinkedIn:

https://www.linkedin.com/in/mt26691/
