# Go Programming for Beginners: Build Real Backend Services

This repository contains the source code for the [Go Programming for Beginners: Build Real Backend Services](https://dalabs.academy) course.

## Prerequisites

- [Go](https://go.dev/dl/) 1.22 or newer (this repo was built with Go 1.26)
- [golangci-lint](https://golangci-lint.run/) (install with `brew install golangci-lint`)
- Comfort using a terminal and a code editor (VS Code with the Go extension recommended)
- [Git](https://git-scm.com/) installed

## Start Branch

```bash
git checkout 08-control-flow-start
```

The start branch adds a new `control_flow.go` file with three demo functions — `loopsDemo`, `fizzBuzzDemo`, and `ifAndSwitchDemo` — whose bodies are stubbed out with `// TODO` comments. `main.go` prints the `== Control Flow ==` header and calls each demo. It still compiles and runs (`make run`), printing just the section headers, ready for you to fill in the loops, FizzBuzz, `if`, and `switch` code.

## Finish Branch

```bash
git checkout 08-control-flow-finish
```

The finish branch completes the control-flow examples: a three-clause `for` loop and a `for range` over a slice of task titles, a FizzBuzz loop (1–15) built on a condition-less `switch` with multiple-value cases, an `if avg, err := mathx.Average(...); err != nil` error check, and a `switch` on a status value with multiple values per case. `make run` prints real output.

## Lesson

[View the lesson on dalabs.academy](<!-- dalabs:08-control-flow -->)

## Running the Program

```bash
make run     # go run .   -> runs the program
make build   # go build   -> compiles the binary
make fmt     # go fmt ./... -> formats the code
make vet     # go vet ./... -> reports suspicious code
make lint    # golangci-lint run -> runs the linter
make help    # list the available targets
```

On the finish branch, `make run` prints real control-flow output:

```
== Control Flow ==

-- Loops --
count up: 1 2 3 4 5 
task 0: write code
task 1: run tests
task 2: ship it

-- FizzBuzz --
1
2
Fizz
4
Buzz
Fizz
7
8
Fizz
Buzz
11
Fizz
13
14
FizzBuzz

-- if & switch --
average score: 8.0
todo         -> still open
in-progress  -> still open
done         -> finished
archived     -> unknown
```

> **Note:** The `Makefile` grows later in the course — a `test` target arrives with the testing section and a `migrate` target with database migrations. The same `make lint` runs in CI at the end of the course.

## Contact

If you have any questions, feedback, or just want to connect, feel free to reach out to me on LinkedIn:

https://www.linkedin.com/in/mt26691/
