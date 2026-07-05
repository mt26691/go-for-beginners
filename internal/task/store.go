package task

import "sync"

// Store is an in-memory task store: a map guarded by a mutex because each
// HTTP request runs in its own goroutine. Chapter 23 extracts an interface
// from it so a PostgreSQL-backed store can take its place later.
type Store struct {
	mu    sync.Mutex
	tasks map[int]Task
}

func NewStore() *Store {
	return &Store{tasks: make(map[int]Task)}
}

func (s *Store) List() []Task {
	s.mu.Lock()
	defer s.mu.Unlock()

	tasks := make([]Task, 0, len(s.tasks))
	for _, t := range s.tasks {
		tasks = append(tasks, t)
	}
	return tasks
}
