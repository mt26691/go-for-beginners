# Go Programming for Beginners: Build Real Backend Services

This repository contains the source code for the [Go Programming for Beginners: Build Real Backend Services](https://dalabs.academy) course.

## Prerequisites

- [Go](https://go.dev/dl/) 1.22 or newer (this repo was built with Go 1.26)
- [golangci-lint](https://golangci-lint.run/) (install with `brew install golangci-lint`)
- Comfort using a terminal and a code editor (VS Code with the Go extension recommended)
- [Git](https://git-scm.com/) installed

## Start Branch

```bash
git checkout 05-developer-tooling-start
```

The start branch is the same working `Hello, Go!` program from Chapter 4 — no tooling yet.

## Finish Branch

```bash
git checkout 05-developer-tooling-finish
```

This branch adds developer tooling around the program: a `Makefile` that wraps the commands you type constantly, and a `.golangci.yml` config for `golangci-lint`.

## Lesson

[View the lesson on dalabs.academy](<!-- dalabs:05-developer-tooling -->)

## Using the Makefile

```bash
make run     # go run .   -> prints "Hello, Go!"
make build   # go build   -> compiles the binary
make fmt     # go fmt ./... -> formats the code
make vet     # go vet ./... -> reports suspicious code
make lint    # golangci-lint run -> runs the linter
make help    # list the available targets
```

On the clean Hello-World program, `make fmt`, `make vet`, and `make lint` all pass with no issues:

```
golangci-lint run
0 issues.
```

> **Note:** The `Makefile` grows later in the course — a `test` target arrives with the testing section and a `migrate` target with database migrations. The same `make lint` runs in CI at the end of the course.

## Contact

If you have any questions, feedback, or just want to connect, feel free to reach out to me on LinkedIn:

https://www.linkedin.com/in/mt26691/
