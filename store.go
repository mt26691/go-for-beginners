package main

import (
	"errors"
	"fmt"
	"io"
	"sort"
)

// ErrTaskNotFound is a sentinel error: a single, package-level error value that
// callers compare against with errors.Is. Find returns it when no task has the
// requested ID. Later in the course the HTTP layer maps this one error to a 404.
var ErrTaskNotFound = errors.New("task not found")

// TaskStore describes the behavior a task store must provide, without saying how.
// Any type whose methods match this set satisfies it automatically — Go has no
// "implements" keyword. This is the seam the whole backend hangs on: the
// in-memory store now and the Postgres store later both satisfy this one
// interface, so the code that depends on it never changes.
type TaskStore interface {
	Add(t Task) Task           // assigns an ID, stores the task, and returns it
	Get(id int) (Task, bool)   // comma-ok lookup: the task and whether it was found
	Find(id int) (Task, error) // error-returning lookup: ErrTaskNotFound when missing
	All() []Task               // a snapshot of every task, sorted by ID
}

// InMemoryStore keeps tasks in a map[int]Task keyed by ID. The constructor
// returns the concrete *InMemoryStore (accept interfaces, return structs).
type InMemoryStore struct {
	tasks  map[int]Task
	nextID int
}

// NewInMemoryStore returns a ready-to-use in-memory store.
func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{tasks: make(map[int]Task)}
}

// Add assigns the next ID, stores the task, and returns the stored copy.
func (s *InMemoryStore) Add(t Task) Task {
	s.nextID++
	t.ID = s.nextID
	s.tasks[t.ID] = t
	return t
}

// Get looks a task up by ID using the comma-ok idiom.
func (s *InMemoryStore) Get(id int) (Task, bool) {
	t, ok := s.tasks[id]
	return t, ok
}

// Find looks a task up by ID and returns an error instead of a bool. When the
// id is missing it returns the sentinel ErrTaskNotFound, which callers match
// with errors.Is and the HTTP layer later maps to a 404.
func (s *InMemoryStore) Find(id int) (Task, error) {
	t, ok := s.tasks[id]
	if !ok {
		return Task{}, ErrTaskNotFound
	}
	return t, nil
}

// All returns every task sorted by ID so the output is deterministic. Map
// iteration order is unspecified, so we sort before returning.
func (s *InMemoryStore) All() []Task {
	out := make([]Task, 0, len(s.tasks))
	for _, t := range s.tasks {
		out = append(out, t)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].ID < out[j].ID })
	return out
}

// LoggingStore wraps another TaskStore and logs every call before delegating.
// It satisfies TaskStore too, so it can stand in anywhere a TaskStore is wanted.
type LoggingStore struct {
	inner TaskStore
	out   io.Writer
}

// NewLoggingStore wraps inner and writes its log lines to out.
func NewLoggingStore(inner TaskStore, out io.Writer) *LoggingStore {
	return &LoggingStore{inner: inner, out: out}
}

// Add logs the call, then delegates to the wrapped store. Writing to s.out can
// fail, so we discard the error with _ — fine for a demo log, but Chapter 12
// shows when you must check it.
func (s *LoggingStore) Add(t Task) Task {
	_, _ = fmt.Fprintf(s.out, "  [log] Add(%q)\n", t.Title)
	return s.inner.Add(t)
}

// Get logs the call, then delegates to the wrapped store.
func (s *LoggingStore) Get(id int) (Task, bool) {
	_, _ = fmt.Fprintf(s.out, "  [log] Get(%d)\n", id)
	return s.inner.Get(id)
}

// Find logs the call, then delegates to the wrapped store.
func (s *LoggingStore) Find(id int) (Task, error) {
	_, _ = fmt.Fprintf(s.out, "  [log] Find(%d)\n", id)
	return s.inner.Find(id)
}

// All logs the call, then delegates to the wrapped store.
func (s *LoggingStore) All() []Task {
	_, _ = fmt.Fprintln(s.out, "  [log] All()")
	return s.inner.All()
}

// runStore takes the TaskStore interface, not a concrete type, so it works with
// any implementation. interfacesDemo calls it with both stores unchanged.
func runStore(s TaskStore) {
	s.Add(Task{Title: "write code"})
	s.Add(Task{Title: "ship it"})

	if t, ok := s.Get(1); ok {
		fmt.Println("  Get(1):", t)
	}
	if _, ok := s.Get(99); !ok {
		fmt.Println("  Get(99): not found")
	}

	fmt.Println("  All():")
	for _, t := range s.All() {
		fmt.Println("   ", t)
	}
}
