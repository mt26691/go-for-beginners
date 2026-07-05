package task

import (
	"context"
	"sync"
	"time"
)

// Store is an in-memory task store: a map guarded by a mutex because each
// HTTP request runs in its own goroutine. Chapter 23 extracts an interface
// from it so a PostgreSQL-backed store can take its place later.
type Store struct {
	mu     sync.Mutex
	tasks  map[int]Task
	nextID int
}

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
