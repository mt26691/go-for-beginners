# Go Programming for Beginners: Build Real Backend Services

This repository contains the source code for the [Go Programming for Beginners: Build Real Backend Services](https://dalabs.academy) course.

## Prerequisites

- [Go](https://go.dev/dl/) 1.22 or newer (this repo was built with Go 1.26)
- [golangci-lint](https://golangci-lint.run/) (install with `brew install golangci-lint`)
- Comfort using a terminal and a code editor (VS Code with the Go extension recommended)
- [Git](https://git-scm.com/) installed

## Start Branch

```bash
git checkout 10-slices-and-maps-start
```

The start branch adds `collections.go` with three demo functions — `slicesDemo`, `mapsDemo`, and `nilMapDemo` — whose bodies are stubbed out with `// TODO` comments (each prints only its section header). `main.go` keeps the structs demos from Chapter 9 and adds a new `== Slices & Maps ==` section that calls the three new demos. It still compiles and runs (`make run`), printing the structs output followed by the empty `Slices`, `Maps`, and `Nil maps` headers — ready for you to fill in slices (`make`/`append`, `len`/`cap`, slice expressions, ranging, and the aliasing gotcha), maps (literals, the comma-ok idiom, `delete`, sorted-key iteration), and the nil-map read-vs-write trap.

## Finish Branch

```bash
git checkout 10-slices-and-maps-finish
```

The finish branch completes `collections.go`: `slicesDemo` builds a `[]Task` with `make` and `append`, prints `len` and `cap`, takes a slice expression (`tasks[1:3]`), ranges over it, and makes the **aliasing** trap visible — two slices that share a backing array, where a write through one shows up through the other. `mapsDemo` builds a `map[int]Task` keyed by ID (the shape of the in-memory store you build in Chapter 20), looks tasks up with the **comma-ok** idiom for a present and an absent key, `delete`s one, and iterates in **sorted-key order** because map iteration order is unspecified. `nilMapDemo` shows that reading a nil map returns the zero value safely (a write would panic). `make run` prints real, deterministic output.

## Lesson

[View the lesson on dalabs.academy](<!-- dalabs:10-slices-and-maps -->)

## Running the Program

```bash
make run     # go run .   -> runs the program
make build   # go build   -> compiles the binary
make fmt     # go fmt ./... -> formats the code
make vet     # go vet ./... -> reports suspicious code
make lint    # golangci-lint run -> runs the linter
make help    # list the available targets
```

On the finish branch, `make run` prints real, deterministic output (the Chapter 9 structs section, then the new slices and maps section):

```
== Slices & Maps ==

-- Slices --
len=3 cap=4 after three appends
  tasks[0] = #1 "write code" [open]
  tasks[1] = #2 "run tests" [open]
  tasks[2] = #3 "ship it" [open]
middle := tasks[1:3] -> len=2 cap=3
after middle[0].Title = ...:
  tasks[1] = #2 "RUN TESTS (edited via middle)" [open]
  middle[0] = #2 "RUN TESTS (edited via middle)" [open]

-- Maps --
found id 2: #2 "run tests" [open]
id 99 is not in the map (comma-ok said so)
after delete(byID, 4): 3 tasks remain
tasks in sorted-key order:
  1 -> #1 "write code" [open]
  2 -> #2 "run tests" [open]
  3 -> #3 "ship it" [done]

-- Nil maps --
read from nil map: ok=false, value={ID:0 Title: Done:false}
writing to a nil map would panic; make() the map first
```

> **Note:** The slices output makes the aliasing trap visible — `tasks[1]` and `middle[0]` both change because the slice expression `tasks[1:3]` shares the same backing array. The maps section iterates by **sorted keys** on purpose: Go's map iteration order is unspecified, so ranging in raw order would not be reproducible.

## Contact

If you have any questions, feedback, or just want to connect, feel free to reach out to me on LinkedIn:

https://www.linkedin.com/in/mt26691/
