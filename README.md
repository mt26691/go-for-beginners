# Go Programming for Beginners: Build Real Backend Services

This repository contains the source code for the [Go Programming for Beginners: Build Real Backend Services](https://dalabs.academy) course.

## Prerequisites

- [Go](https://go.dev/dl/) 1.22 or newer (this repo was built with Go 1.26)
- [golangci-lint](https://golangci-lint.run/) (install with `brew install golangci-lint`)
- Comfort using a terminal and a code editor (VS Code with the Go extension recommended)
- [Git](https://git-scm.com/) installed

## Start Branch

```bash
git checkout 06-variables-and-types-start
```

The start branch is a scaffold: `main.go` has four small functions (`variablesDemo`, `zeroValuesDemo`, `basicTypesDemo`, `constantsDemo`) with `// TODO` comments describing the variables, types, and constants you will add. It still compiles and runs — `make run` prints the section headers with empty bodies underneath, ready for you to fill in.

## Finish Branch

```bash
git checkout 06-variables-and-types-finish
```

This branch completes the scaffold: each function declares and prints real variables, shows zero values, demonstrates Go's basic types and type inference, and defines constants with an `iota` enumeration.

## Lesson

[View the lesson on dalabs.academy]({URL})

## Using the Makefile

```bash
make run     # go run .   -> runs the program
make build   # go build   -> compiles the binary
make fmt     # go fmt ./... -> formats the code
make vet     # go vet ./... -> reports suspicious code
make lint    # golangci-lint run -> runs the linter
make help    # list the available targets
```

On the finish branch, `make run` prints labeled output for each concept:

```
== Variables, Types & Constants ==

-- Variables --
language: Go
temperature: 36
language (reassigned): Golang

-- Zero values --
int:     0
float64: 0
string:  ""
bool:    false

-- Basic types --
42        has type int
9000000000 has type int64
19.99     has type float64
"héllo"   has type string
true      has type bool
"héllo": 6 bytes, 5 runes

-- Constants --
pi: 3.14159
levels: debug=0 info=1 warn=2 error=3
```

> **Note:** The `Makefile` grows later in the course — a `test` target arrives with the testing section and a `migrate` target with database migrations. The same `make lint` runs in CI at the end of the course.

## Contact

If you have any questions, feedback, or just want to connect, feel free to reach out to me on LinkedIn:

https://www.linkedin.com/in/mt26691/
