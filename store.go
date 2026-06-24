package main

import (
	"fmt"
	"io"
	"sort"
)

// TaskStore describes the behavior a task store must provide, without saying how.
// Any type whose methods match this set satisfies it automatically — Go has no
// "implements" keyword. This is the seam the whole backend hangs on: the
// in-memory store now and the Postgres store later both satisfy this one
// interface, so the code that depends on it never changes.
type TaskStore interface {
	Add(t Task) Task         // assigns an ID, stores the task, and returns it
	Get(id int) (Task, bool) // comma-ok lookup: the task and whether it was found
	All() []Task             // a snapshot of every task, sorted by ID
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
	// TODO: increment s.nextID, set t.ID to it, store t in s.tasks, return t.
	s.nextID++
	return t
}

// Get looks a task up by ID using the comma-ok idiom.
func (s *InMemoryStore) Get(id int) (Task, bool) {
	// TODO: read s.tasks[id] with comma-ok and return (task, ok).
	return Task{}, false
}

// All returns every task sorted by ID so the output is deterministic.
func (s *InMemoryStore) All() []Task {
	// TODO: collect s.tasks into a slice, sort it by ID with sort.Slice, return it.
	return nil
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

// Add logs the call, then delegates to the wrapped store.
func (s *LoggingStore) Add(t Task) Task {
	// TODO: write a log line to s.out, then return s.inner.Add(t).
	return s.inner.Add(t)
}

// Get logs the call, then delegates to the wrapped store.
func (s *LoggingStore) Get(id int) (Task, bool) {
	// TODO: write a log line to s.out, then return s.inner.Get(id).
	return s.inner.Get(id)
}

// All logs the call, then delegates to the wrapped store.
func (s *LoggingStore) All() []Task {
	// TODO: write a log line to s.out, then return s.inner.All().
	return s.inner.All()
}

// runStore takes the TaskStore interface, not a concrete type, so it works with
// any implementation. We call it with both stores to prove they are swappable.
func runStore(s TaskStore) {
	// TODO: Add two tasks, Get one present and one absent, print s.All().
	_ = s
	_ = sort.Ints
	_ = fmt.Sprint
}
