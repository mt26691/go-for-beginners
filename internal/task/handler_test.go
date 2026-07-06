package task_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/mt26691/go-for-beginners/internal/task"
)

// fakeStore is a test double that satisfies task.TaskStore without a database.
// It keeps tasks in a plain map, so a test can control exactly what storage
// returns — including task.ErrNotFound for an unknown ID. The interface seam
// from Chapter 23 is what lets us inject it in place of the real store.
type fakeStore struct {
	tasks  map[int]task.Task
	nextID int
}

// Compile-time proof that *fakeStore satisfies the interface the service needs.
var _ task.TaskStore = (*fakeStore)(nil)

func newFakeStore() *fakeStore {
	return &fakeStore{tasks: make(map[int]task.Task)}
}

func (f *fakeStore) Create(_ context.Context, t task.Task) (task.Task, error) {
	f.nextID++
	t.ID = f.nextID
	t.CreatedAt = time.Now()
	f.tasks[t.ID] = t
	return t, nil
}

func (f *fakeStore) List(_ context.Context) ([]task.Task, error) {
	tasks := make([]task.Task, 0, len(f.tasks))
	for _, t := range f.tasks {
		tasks = append(tasks, t)
	}
	return tasks, nil
}

func (f *fakeStore) Get(_ context.Context, id int) (task.Task, error) {
	t, ok := f.tasks[id]
	if !ok {
		return task.Task{}, task.ErrNotFound
	}
	return t, nil
}

func (f *fakeStore) Update(_ context.Context, id int, t task.Task) (task.Task, error) {
	existing, ok := f.tasks[id]
	if !ok {
		return task.Task{}, task.ErrNotFound
	}
	t.ID = existing.ID
	t.CreatedAt = existing.CreatedAt
	f.tasks[id] = t
	return t, nil
}

func (f *fakeStore) Delete(_ context.Context, id int) error {
	if _, ok := f.tasks[id]; !ok {
		return task.ErrNotFound
	}
	delete(f.tasks, id)
	return nil
}

// newTestServer wires a fresh fake store through the real service and handler
// and returns the mux plus the store. Each test gets its own isolated storage,
// so state never leaks between cases.
func newTestServer() (http.Handler, *fakeStore) {
	store := newFakeStore()
	svc := task.NewService(store)
	h := task.NewHandler(svc)
	mux := http.NewServeMux()
	h.Routes(mux)
	return mux, store
}

// jsonRequest builds a request with a JSON-encoded body and the right header.
func jsonRequest(method, target string, body any) *http.Request {
	var buf bytes.Buffer
	if body != nil {
		if err := json.NewEncoder(&buf).Encode(body); err != nil {
			panic(err)
		}
	}
	req := httptest.NewRequest(method, target, &buf)
	req.Header.Set("Content-Type", "application/json")
	return req
}

// decodeInto reads the recorder's JSON body into dst.
func decodeInto(t *testing.T, rec *httptest.ResponseRecorder, dst any) {
	t.Helper()
	if err := json.NewDecoder(rec.Body).Decode(dst); err != nil {
		t.Fatalf("decode response body: %v", err)
	}
}

func TestCreateTask(t *testing.T) {
	t.Run("valid task returns 201 and the created task", func(t *testing.T) {
		srv, _ := newTestServer()

		req := jsonRequest(http.MethodPost, "/tasks", map[string]string{"title": "Buy milk"})
		rec := httptest.NewRecorder()
		srv.ServeHTTP(rec, req)

		if rec.Code != http.StatusCreated {
			t.Fatalf("status = %d, want %d", rec.Code, http.StatusCreated)
		}
		if got := rec.Header().Get("Content-Type"); got != "application/json" {
			t.Errorf("Content-Type = %q, want application/json", got)
		}

		var got task.Task
		decodeInto(t, rec, &got)
		if got.Title != "Buy milk" {
			t.Errorf("title = %q, want %q", got.Title, "Buy milk")
		}
		if got.ID == 0 {
			t.Error("id = 0, want a server-assigned id")
		}
	})

	t.Run("empty title returns 400 and an error envelope", func(t *testing.T) {
		srv, _ := newTestServer()

		req := jsonRequest(http.MethodPost, "/tasks", map[string]string{"title": ""})
		rec := httptest.NewRecorder()
		srv.ServeHTTP(rec, req)

		if rec.Code != http.StatusBadRequest {
			t.Fatalf("status = %d, want %d", rec.Code, http.StatusBadRequest)
		}

		var got map[string]string
		decodeInto(t, rec, &got)
		if got["error"] != "title is required" {
			t.Errorf("error = %q, want %q", got["error"], "title is required")
		}
	})
}

func TestListTasks(t *testing.T) {
	srv, store := newTestServer()
	if _, err := store.Create(context.Background(), task.Task{Title: "Buy milk"}); err != nil {
		t.Fatalf("seed store: %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/tasks", nil)
	rec := httptest.NewRecorder()
	srv.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	var got []task.Task
	decodeInto(t, rec, &got)
	if len(got) != 1 {
		t.Fatalf("len(tasks) = %d, want 1", len(got))
	}
	if got[0].Title != "Buy milk" {
		t.Errorf("title = %q, want %q", got[0].Title, "Buy milk")
	}
}

func TestGetTask(t *testing.T) {
	t.Run("existing id returns 200 and the task", func(t *testing.T) {
		srv, store := newTestServer()
		created, err := store.Create(context.Background(), task.Task{Title: "Buy milk"})
		if err != nil {
			t.Fatalf("seed store: %v", err)
		}

		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/tasks/%d", created.ID), nil)
		rec := httptest.NewRecorder()
		srv.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
		}

		var got task.Task
		decodeInto(t, rec, &got)
		if got.ID != created.ID {
			t.Errorf("id = %d, want %d", got.ID, created.ID)
		}
	})

	t.Run("missing id returns 404", func(t *testing.T) {
		srv, _ := newTestServer()

		req := httptest.NewRequest(http.MethodGet, "/tasks/999", nil)
		rec := httptest.NewRecorder()
		srv.ServeHTTP(rec, req)

		if rec.Code != http.StatusNotFound {
			t.Fatalf("status = %d, want %d", rec.Code, http.StatusNotFound)
		}

		var got map[string]string
		decodeInto(t, rec, &got)
		if got["error"] != "task not found" {
			t.Errorf("error = %q, want %q", got["error"], "task not found")
		}
	})
}
