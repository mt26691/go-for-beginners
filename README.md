# Go Programming for Beginners: Build Real Backend Services

This repository contains the source code for the [Go Programming for Beginners: Build Real Backend Services](https://dalabs.academy) course.

## Prerequisites

- [Go](https://go.dev/dl/) 1.22 or newer (this repo was built with Go 1.26)
- [golangci-lint](https://golangci-lint.run/) (install with `brew install golangci-lint`)
- Comfort using a terminal and a code editor (VS Code with the Go extension recommended)
- [Git](https://git-scm.com/) installed

## Start Branch

```bash
git checkout 07-functions-and-packages-start
```

The start branch sets up two new packages of our own — `mathx` and `textx` — each with exported functions stubbed out with `// TODO` comments. `main.go` prints the section headers with empty bodies, ready for you to fill in. It still compiles and runs (`make run`).

## Finish Branch

```bash
git checkout 07-functions-and-packages-finish
```

The finish branch completes the packages: `mathx.MinMax` (variadic, named returns), `mathx.Average` (returns a value and an error), and `textx.Label` (uses the standard library plus the third-party `golang.org/x/text` package). `main.go` imports both packages by their full module path and prints real results.

## Lesson

[View the lesson on dalabs.academy](<!-- dalabs:07-functions-and-packages -->)

## Running the Program

```bash
make run     # go run .   -> runs the program
make build   # go build   -> compiles the binary
make fmt     # go fmt ./... -> formats the code
make vet     # go vet ./... -> reports suspicious code
make lint    # golangci-lint run -> runs the linter
make help    # list the available targets
```

On the start branch, `make run` prints the section headers with empty bodies:

```
== Functions & Packages ==

-- Functions --

-- Packages --
```

> **Note:** The `Makefile` grows later in the course — a `test` target arrives with the testing section and a `migrate` target with database migrations. The same `make lint` runs in CI at the end of the course.

## Contact

If you have any questions, feedback, or just want to connect, feel free to reach out to me on LinkedIn:

https://www.linkedin.com/in/mt26691/
