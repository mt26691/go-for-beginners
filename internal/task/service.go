package task

import "context"

// TaskStore is the behavior the service needs from storage, and nothing more.
// The in-memory Store satisfies it today (see the compile-time check in
// store.go); a PostgreSQL-backed store will satisfy the same interface in
// Section 6, so neither the service nor the handlers change when storage does.
type TaskStore interface {
	Create(ctx context.Context, t Task) (Task, error)
	List(ctx context.Context) ([]Task, error)
	Get(ctx context.Context, id int) (Task, error)
	Update(ctx context.Context, id int, t Task) (Task, error)
	Delete(ctx context.Context, id int) error
}

// Service holds the business logic that sits between the HTTP handlers and
// storage. It depends on the TaskStore interface, not the concrete Store, so
// the handlers stay ignorant of how tasks are actually stored.
type Service struct {
	store TaskStore
}

// NewService injects the store the service will use. The dependency arrives
// through the constructor rather than being created inside, so a test can pass
// a fake store and main can pass the real one.
func NewService(store TaskStore) *Service {
	return &Service{store: store}
}

// Create validates the task, then hands it to storage. Validation lives here
// so every path that creates a task goes through the same rules.
func (s *Service) Create(ctx context.Context, t Task) (Task, error) {
	if err := t.Validate(); err != nil {
		return Task{}, err
	}
	return s.store.Create(ctx, t)
}

func (s *Service) List(ctx context.Context) ([]Task, error) {
	return s.store.List(ctx)
}

func (s *Service) Get(ctx context.Context, id int) (Task, error) {
	return s.store.Get(ctx, id)
}

// Update validates the incoming task before replacing the stored one.
func (s *Service) Update(ctx context.Context, id int, t Task) (Task, error) {
	if err := t.Validate(); err != nil {
		return Task{}, err
	}
	return s.store.Update(ctx, id, t)
}

func (s *Service) Delete(ctx context.Context, id int) error {
	return s.store.Delete(ctx, id)
}
