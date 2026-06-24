# Go Programming for Beginners: Build Real Backend Services

This repository contains the source code for the [Go Programming for Beginners: Build Real Backend Services](https://dalabs.academy) course.

## Prerequisites

- [Go](https://go.dev/dl/) 1.22 or newer (this repo was built with Go 1.26)
- [golangci-lint](https://golangci-lint.run/) (install with `brew install golangci-lint`)
- Comfort using a terminal and a code editor (VS Code with the Go extension recommended)
- [Git](https://git-scm.com/) installed

## Start Branch

```bash
git checkout 11-interfaces-start
```

The start branch adds two new files. `store.go` defines the `TaskStore` interface (`Add`, `Get`, `All`) and two types that will satisfy it — an `InMemoryStore` backed by a `map[int]Task` and a `LoggingStore` that wraps another `TaskStore` — but their method bodies are stubbed with `// TODO` comments (they return zero values or just delegate). `interfaces.go` defines `interfacesDemo` and a `describe(any)` helper, also stubbed. `task.go` gains a `String()` method so `Task` satisfies `fmt.Stringer`. `main.go` adds a new `== Interfaces ==` section that calls `interfacesDemo`. It compiles, vets, lints, and runs (`make run`) — printing the structs and slices/maps output, then the `Interfaces` header — ready for you to fill in the two stores, `runStore`, the type switch in `describe`, and the empty-interface example.

## Finish Branch

```bash
git checkout 11-interfaces-finish
```

The finish branch completes `store.go` and `interfaces.go`. `InMemoryStore` assigns IDs and stores tasks in its map; its `All()` returns a snapshot **sorted by ID** so the output is deterministic. `LoggingStore` logs each call to an `io.Writer`, then delegates to the wrapped store — proving a wrapper can satisfy the same interface. `runStore(s TaskStore)` takes the **interface**, not a concrete type, and is called with both stores unchanged (the swappable seam this whole backend hangs on). `interfacesDemo` also shows a **type assertion**, a **type switch** (in `describe`), and the empty interface `any`. `Task` satisfies `fmt.Stringer` via `String()`. `make run` prints real, deterministic output.

## Lesson

[View the lesson on dalabs.academy](https://dalabs.academy/courses/go-programming-for-beginners-build-real-backend-services/go-language-foundations/interfaces)

## Running the Program

```bash
make run     # go run .   -> runs the program
make build   # go build   -> compiles the binary
make fmt     # go fmt ./... -> formats the code
make vet     # go vet ./... -> reports suspicious code
make lint    # golangci-lint run -> runs the linter
make help    # list the available targets
```

> **Note:** On the finish branch `make run` prints real, deterministic output. The `Interfaces` section runs the same `runStore(s TaskStore)` function against an `InMemoryStore` and a `LoggingStore` and gets the same results — the swappable seam. `All()` sorts by ID so the output never changes between runs, and `LoggingStore` interleaves its `[log]` lines to show the wrapper satisfying the same interface.

## Contact

If you have any questions, feedback, or just want to connect, feel free to reach out to me on LinkedIn:

https://www.linkedin.com/in/mt26691/
