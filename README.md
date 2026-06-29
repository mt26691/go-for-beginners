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

The start branch carries over the working "Hello from Go!" server from Chapter 14. `main.go` has `helloHandler` registered on `"/"` and `http.ListenAndServe(":8080", nil)`. You will build on this by adding `greetHandler`, `echoHandler`, and a custom `pingHandler` type to explore the full request/response API.

## Finish Branch

```bash
git checkout 15-handlers-and-responses-finish
```

The finish branch adds three new handlers on top of the Chapter 14 server. `greetHandler` reads a query parameter and returns a 400 when it is missing. `echoHandler` reads the request body with `io.ReadAll`, sets a `Content-Type` header, and returns 201. `pingHandler` is a custom struct that satisfies `http.Handler` via `ServeHTTP`, demonstrating the interface directly rather than the `http.HandlerFunc` adapter.

## Lesson

[View the lesson on dalabs.academy](<!-- dalabs:15-handlers-and-responses -->)
<!-- After publishing, the /publish-chapter skill replaces the placeholder above with the actual URL -->

## Running the Server

```bash
go run .
```

Then in another terminal:

```bash
curl http://localhost:8080/
# Hello from Go!
```

```bash
make build   # go build   -> compiles the binary
make fmt     # go fmt ./... -> formats the code
make vet     # go vet ./... -> reports suspicious code
make lint    # golangci-lint run -> runs the linter
make help    # list the available targets
```

> **Note:** This is the start branch — `main.go` still contains only the Chapter 14 "Hello from Go!" server. Your goal is to add `greetHandler`, `echoHandler`, and `pingHandler`, then wire them to the mux to reach the finish state.

## Contact

If you have any questions, feedback, or just want to connect, feel free to reach out to me on LinkedIn:

https://www.linkedin.com/in/mt26691/
