# Go Programming for Beginners: Build Real Backend Services

This repository contains the source code for the [Go Programming for Beginners: Build Real Backend Services](https://dalabs.academy) course.

## Prerequisites

- [Go](https://go.dev/dl/) 1.22 or newer (this repo was built with Go 1.26)
- [golangci-lint](https://golangci-lint.run/) (install with `brew install golangci-lint`)
- Comfort using a terminal and a code editor (VS Code with the Go extension recommended)
- [Git](https://git-scm.com/) installed

## Start Branch

```bash
git checkout 14-first-http-server-start
```

The start branch is a clean slate for Section 3. All Section 2 demo files (`collections.go`, `concurrency.go`, `errorsdemo.go`, `interfaces.go`, `structs.go`, `store.go`, `task.go`) and the `mathx/` and `textx/` helper packages have been removed. `main.go` contains an empty `func main() {}` — ready for you to wire up your first HTTP server using `net/http`.

## Finish Branch

```bash
git checkout 14-first-http-server-finish
```

The finish branch adds a minimal, working HTTP server in `main.go`. A single handler is registered for `"/"` via `http.HandleFunc`, writing a plain-text greeting to the response. `log.Fatal(http.ListenAndServe(":8080", nil))` starts the server and surfaces any error instead of silently ignoring it. Run `go run .` and hit `http://localhost:8080` with a browser or `curl` to see the response.

## Lesson

[View the lesson on dalabs.academy](<!-- dalabs:14-first-http-server -->)
<!-- After publishing, the /publish-chapter skill replaces the placeholder above with the actual URL -->

## Running the Server

```bash
go run .
```

Then in another terminal:

```bash
curl http://localhost:8080
# Hello from Go!
```

```bash
make build   # go build   -> compiles the binary
make fmt     # go fmt ./... -> formats the code
make vet     # go vet ./... -> reports suspicious code
make lint    # golangci-lint run -> runs the linter
make help    # list the available targets
```

> **Note:** This is the finish branch — `main.go` contains a working HTTP server. `go run .` starts it on port 8080. Use `curl http://localhost:8080` or open your browser to see `Hello from Go!`. Stop the server with `Ctrl+C`.

## Contact

If you have any questions, feedback, or just want to connect, feel free to reach out to me on LinkedIn:

https://www.linkedin.com/in/mt26691/
