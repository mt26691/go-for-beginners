package task_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
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
// so state never leaks between cases — which is what makes t.Parallel() safe.
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

// TestCreateTask drives POST /tasks from a table of cases. Each row is one
// request/response scenario, and the loop turns every row into a named subtest
// with t.Run, so a failure points at the exact case. wantError is empty when the
// case should succeed and non-empty when it should return the JSON error
// envelope {"error": ...}.
func TestCreateTask(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name       string
		title      string
		wantStatus int
		wantError  string
	}{
		{
			name:       "valid title returns 201",
			title:      "Buy milk",
			wantStatus: http.StatusCreated,
		},
		{
			name:       "empty title returns 400",
			title:      "",
			wantStatus: http.StatusBadRequest,
			wantError:  "title is required",
		},
		{
			name:       "too-long title returns 400",
			title:      strings.Repeat("a", task.MaxTitleLength+1),
			wantStatus: http.StatusBadRequest,
			wantError:  "title must be 200 characters or fewer",
		},
	}

	for _, tc := range cases {
		// Since Go 1.22 each iteration gets a fresh tc, so the old `tc := tc`
		// copy is no longer needed to make t.Parallel() capture the right case.
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			srv, _ := newTestServer()
			rec := httptest.NewRecorder()
			srv.ServeHTTP(rec, jsonRequest(http.MethodPost, "/tasks", map[string]string{"title": tc.title}))

			if rec.Code != tc.wantStatus {
				t.Fatalf("status = %d, want %d", rec.Code, tc.wantStatus)
			}

			if tc.wantError != "" {
				var got map[string]string
				decodeInto(t, rec, &got)
				if got["error"] != tc.wantError {
					t.Errorf("error = %q, want %q", got["error"], tc.wantError)
				}
				return
			}

			var got task.Task
			decodeInto(t, rec, &got)
			if got.Title != tc.title {
				t.Errorf("title = %q, want %q", got.Title, tc.title)
			}
			if got.ID == 0 {
				t.Error("id = 0, want a server-assigned id")
			}
		})
	}
}

// TestGetTask drives GET /tasks/{id} from its own small table. The id cases —
// found, non-numeric, and missing — do not fit the create table, so they get a
// second table instead of being forced into one giant one. When seed is true the
// case creates a task first and requests its real id; otherwise it uses target.
func TestGetTask(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name       string
		seed       bool
		target     string
		wantStatus int
		wantError  string
	}{
		{
			name:       "existing id returns 200",
			seed:       true,
			wantStatus: http.StatusOK,
		},
		{
			name:       "non-numeric id returns 400",
			target:     "/tasks/abc",
			wantStatus: http.StatusBadRequest,
			wantError:  "invalid task id",
		},
		{
			name:       "missing id returns 404",
			target:     "/tasks/999",
			wantStatus: http.StatusNotFound,
			wantError:  "task not found",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			srv, store := newTestServer()

			target := tc.target
			if tc.seed {
				created, err := store.Create(context.Background(), task.Task{Title: "Buy milk"})
				if err != nil {
					t.Fatalf("seed store: %v", err)
				}
				target = fmt.Sprintf("/tasks/%d", created.ID)
			}

			rec := httptest.NewRecorder()
			srv.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, target, nil))

			if rec.Code != tc.wantStatus {
				t.Fatalf("status = %d, want %d", rec.Code, tc.wantStatus)
			}

			if tc.wantError != "" {
				var got map[string]string
				decodeInto(t, rec, &got)
				if got["error"] != tc.wantError {
					t.Errorf("error = %q, want %q", got["error"], tc.wantError)
				}
				return
			}

			var got task.Task
			decodeInto(t, rec, &got)
			if got.Title != "Buy milk" {
				t.Errorf("title = %q, want %q", got.Title, "Buy milk")
			}
		})
	}
}

// TestListTasks stays a single straight-line test. It is one scenario, not a
// family of similar cases, so a table would only add ceremony — a reminder that
// the table pattern is for many similar cases, not every test.
func TestListTasks(t *testing.T) {
	t.Parallel()

	srv, store := newTestServer()
	if _, err := store.Create(context.Background(), task.Task{Title: "Buy milk"}); err != nil {
		t.Fatalf("seed store: %v", err)
	}

	rec := httptest.NewRecorder()
	srv.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/tasks", nil))

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
