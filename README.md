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

The start branch is the finished Chapter 15 server. `main.go` registers `helloHandler`, `greetHandler`, `echoHandler`, and `pingHandler` on the **default mux** with bare path strings (`http.HandleFunc("/greet", ...)` and `http.Handle("/ping", ...)`), and `http.ListenAndServe(":8080", nil)` uses that default mux. You will rebuild the routing on an explicit `http.NewServeMux()` and use Go 1.22 method-plus-path patterns to map the Task API endpoints.

## Finish Branch

```bash
git checkout 16-routing-with-servemux-finish
```

The finish branch creates an explicit `http.NewServeMux()` and registers method-plus-path patterns: `GET /ping`, `GET /tasks`, `POST /tasks`, and `GET /tasks/{id}` (read with `r.PathValue("id")`). The task handlers are curl-able stubs for now, and `http.ListenAndServe` receives `mux` instead of `nil`.

## Lesson

[View the lesson on dalabs.academy](<!-- dalabs:16-routing-with-servemux -->)
<!-- After publishing, the /publish-chapter skill replaces the placeholder above with the actual URL -->

## Running the Server

```bash
go run .
```

Then in another terminal:

```bash
curl http://localhost:8080/
# Hello from Go!

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

> **Note:** This is the start branch — `main.go` still routes through the default mux with bare path strings, exactly as Chapter 15 left it. Your goal is to switch to an explicit `http.NewServeMux()` with Go 1.22 method-plus-path patterns and route the Task API onto it.

## Contact

If you have any questions, feedback, or just want to connect, feel free to reach out to me on LinkedIn:

https://www.linkedin.com/in/mt26691/
