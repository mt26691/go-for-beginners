# Go Programming for Beginners: Build Real Backend Services

This repository contains the source code for the [Go Programming for Beginners: Build Real Backend Services](https://dalabs.academy) course.

## Prerequisites

- [Go](https://go.dev/dl/) 1.22 or newer (this repo was built with Go 1.26)
- [golangci-lint](https://golangci-lint.run/) (install with `brew install golangci-lint`)
- Comfort using a terminal and a code editor (VS Code with the Go extension recommended)
- [Git](https://git-scm.com/) installed

## Start Branch

```bash
git checkout 12-errors-start
```

The start branch sets the stage for Go's error model. `store.go` gains a sentinel error, `var ErrTaskNotFound = errors.New("task not found")`, and a new `Find(id int) (Task, error)` method is added to the `TaskStore` interface and both stores — but the `InMemoryStore.Find` body is stubbed with a `// TODO`. A new `errorsdemo.go` declares a custom `ValidationError` type, an `errorsDemo()` function, and three helpers (`loadTask`, `validateTitle`, `safeDivide`), all stubbed with `// TODO` comments that return `nil` or zero values. `main.go` adds a new `== Errors ==` section that calls `errorsDemo`. It compiles, vets, lints, and runs (`make run`) — printing the earlier sections, then the `Errors` header with placeholder `<nil>` lines — ready for you to fill in the sentinel lookup, the `%w` wrap, `errors.Is`/`errors.As`, and the `recover` example.

## Finish Branch

```bash
git checkout 12-errors-finish
```

The finish branch completes the error model. `InMemoryStore.Find` returns the task when present and `ErrTaskNotFound` when it is missing — the sentinel the HTTP layer maps to a 404 later in the course. `loadTask` wraps `ErrTaskNotFound` with `fmt.Errorf("...: %w", ...)` so callers can still unwrap it; `errorsDemo` prints the wrapped chain and uses `errors.Is` to find the sentinel through the wrap. `validateTitle` returns a `*ValidationError`, recovered at the call site with `errors.As`. `safeDivide` uses `defer`/`recover` to turn a divide-by-zero panic into a returned error, so the program stays green. `make run` prints real, deterministic output.

## Lesson

[View the lesson on dalabs.academy](https://dalabs.academy/courses/go-programming-for-beginners-build-real-backend-services/go-language-foundations/errors)

## Running the Program

```bash
make run     # go run .   -> runs the program
make build   # go build   -> compiles the binary
make fmt     # go fmt ./... -> formats the code
make vet     # go vet ./... -> reports suspicious code
make lint    # golangci-lint run -> runs the linter
make help    # list the available targets
```

> **Note:** On the start branch the `== Errors ==` section prints placeholder `<nil>` lines because the helper bodies are stubbed. Fill in the `// TODO`s (or check out the finish branch) to see the wrapped error, `errors.Is` matching `ErrTaskNotFound` through a wrap, `errors.As` extracting the `ValidationError`, and `recover` turning a panic into an error.

## Contact

If you have any questions, feedback, or just want to connect, feel free to reach out to me on LinkedIn:

https://www.linkedin.com/in/mt26691/
