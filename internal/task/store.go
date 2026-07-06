package task

import (
	"context"
	"errors"
	"sync"
	"time"
)

// ErrNotFound is returned by the store when no task exists for the given ID.
// Handlers use errors.Is(err, ErrNotFound) to map it to a 404 response.
var ErrNotFound = errors.New("task not found")

// Store is an in-memory task store: a map guarded by a mutex because each
// HTTP request runs in its own goroutine. It satisfies the TaskStore
// interface (see the compile-time check below), so a PostgreSQL-backed store
// can take its place later without touching the service or the handlers.
type Store struct {
	mu     sync.Mutex
	tasks  map[int]Task
	nextID int
}

// Compile-time assertion that *Store satisfies TaskStore. If a method ever
// drifts out of sync with the interface, the build fails here instead of at
// the call site.
var _ TaskStore = (*Store)(nil)

func NewStore() *Store {
	return &Store{tasks: make(map[int]Task)}
}

// Create assigns a server-side ID and creation time, stores the task, and
// returns the stored copy. The context is unused by the in-memory store but
// is part of the signature the database-backed store (Section 6) will need.
func (s *Store) Create(_ context.Context, t Task) (Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.nextID++
	t.ID = s.nextID
	t.CreatedAt = time.Now()
	s.tasks[t.ID] = t
	return t, nil
}

func (s *Store) List(_ context.Context) ([]Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	tasks := make([]Task, 0, len(s.tasks))
	for _, t := range s.tasks {
		tasks = append(tasks, t)
	}
	return tasks, nil
}

// Get returns the task with the given ID, or ErrNotFound if it does not exist.
func (s *Store) Get(_ context.Context, id int) (Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	t, ok := s.tasks[id]
	if !ok {
		return Task{}, ErrNotFound
	}
	return t, nil
}

// Update replaces the stored task with the given ID. It is a full replace: the
// original ID and creation time are preserved, and every other field is taken
// from t. It returns ErrNotFound if no task has that ID.
func (s *Store) Update(_ context.Context, id int, t Task) (Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	existing, ok := s.tasks[id]
	if !ok {
		return Task{}, ErrNotFound
	}

	t.ID = existing.ID
	t.CreatedAt = existing.CreatedAt
	s.tasks[id] = t
	return t, nil
}

// Delete removes the task with the given ID, or returns ErrNotFound if there is
// no task with that ID.
func (s *Store) Delete(_ context.Context, id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.tasks[id]; !ok {
		return ErrNotFound
	}

	delete(s.tasks, id)
	return nil
}
