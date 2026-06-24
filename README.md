# Go Programming for Beginners: Build Real Backend Services

This repository contains the source code for the [Go Programming for Beginners: Build Real Backend Services](https://dalabs.academy) course.

## Prerequisites

- [Go](https://go.dev/dl/) 1.22 or newer (this repo was built with Go 1.26)
- [golangci-lint](https://golangci-lint.run/) (install with `brew install golangci-lint`)
- Comfort using a terminal and a code editor (VS Code with the Go extension recommended)
- [Git](https://git-scm.com/) installed

## Start Branch

```bash
git checkout 09-structs-and-methods-start
```

The start branch adds two new files — `task.go` (defining the `Task` struct with `ID`, `Title`, and `Done`) and `structs.go` (the `structsDemo`, `pointersDemo`, and `embeddingDemo` functions) — with their method bodies and demo logic stubbed out with `// TODO` comments. `main.go` prints the `== Structs, Methods & Pointers ==` header and calls each demo. It still compiles and runs (`make run`), printing just the section headers, ready for you to fill in the struct, the value-receiver vs pointer-receiver methods, the pointer basics, and the embedding example.

## Finish Branch

```bash
git checkout 09-structs-and-methods-finish
```

The finish branch completes the examples: a `Task` struct created three ways (a struct literal, `&Task{}`, and `new`), a read-only `Summary()` method on a **value receiver**, a `markDoneByValue` method that mutates a **copy** (so the change never sticks) contrasted with a `MarkDone()` method on a **pointer receiver** (so the change does stick), a pointer demo using `&` and `*` plus a nil `*Task`, and a `labeledTask` that **embeds** `Task` and calls its promoted method and field. `make run` prints real output that makes the value-vs-pointer contrast visible.

## Lesson

[View the lesson on dalabs.academy](<!-- dalabs:09-structs-and-methods -->)

## Running the Program

```bash
make run     # go run .   -> runs the program
make build   # go build   -> compiles the binary
make fmt     # go fmt ./... -> formats the code
make vet     # go vet ./... -> reports suspicious code
make lint    # golangci-lint run -> runs the linter
make help    # list the available targets
```

On the finish branch, `make run` prints real output showing the value-vs-pointer contrast:

```
== Structs, Methods & Pointers ==

-- Structs --
#1 "write code" [open]
#2 "run tests" [open]
#3 "ship it" [open]

-- Value vs Pointer Receiver --
markDoneByValue set the copy to true, but...
after markDoneByValue: #1 "write code" [open]
after MarkDone:        #1 "write code" [done]

-- Pointers --
count is now 42 (changed through the pointer)
a nil *Task prints as: <nil>

-- Embedding --
#4 "review PR" [done] (label: urgent)
promoted field Title: review PR
```

> **Note:** The `Makefile` grows later in the course — a `test` target arrives with the testing section and a `migrate` target with database migrations. The same `make lint` runs in CI at the end of the course.

## Contact

If you have any questions, feedback, or just want to connect, feel free to reach out to me on LinkedIn:

https://www.linkedin.com/in/mt26691/
