// main.go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	Done        bool   `json:"done"`
	Internal    string `json:"-"`
}

type ctxKey string

const requestIDKey ctxKey = "requestID"

const storeQueryDelay = 500 * time.Millisecond

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Printf("writeJSON: %v", err)
	}
}

func findTask(ctx context.Context, id int) (Task, error) {
	select {
	case <-time.After(storeQueryDelay):
		return Task{ID: id, Title: "Sample task"}, nil
	case <-ctx.Done():
		return Task{}, ctx.Err()
	}
}

func requestIDFrom(ctx context.Context) string {
	if id, ok := ctx.Value(requestIDKey).(string); ok {
		return id
	}
	return "unknown"
}

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "pong")
}

func listTasks(w http.ResponseWriter, r *http.Request) {
	tasks := []Task{
		{ID: 1, Title: "Write the JSON chapter", Done: true, Internal: "seed"},
		{ID: 2, Title: "Record the demo", Description: "curl every route", Done: false},
	}
	writeJSON(w, http.StatusOK, tasks)
}

func getTask(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid task id", http.StatusBadRequest)
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()
	task, err := findTask(ctx, id)
	if err != nil {
		log.Printf("getTask: %v", err)
		http.Error(w, err.Error(), http.StatusGatewayTimeout)
		return
	}
	writeJSON(w, http.StatusOK, task)
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var t Task
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, "invalid JSON body", http.StatusBadRequest)
		return
	}
	writeJSON(w, http.StatusCreated, t)
}

func slowHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	select {
	case <-time.After(5 * time.Second):
		writeJSON(w, http.StatusOK, map[string]string{"status": "done"})
	case <-ctx.Done():
		log.Printf("slow: request cancelled: %v", ctx.Err())
	}
}

func slowStoreHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 100*time.Millisecond)
	defer cancel()
	task, err := findTask(ctx, 1)
	if err != nil {
		log.Printf("slow-store: %v", err)
		http.Error(w, err.Error(), http.StatusGatewayTimeout)
		return
	}
	writeJSON(w, http.StatusOK, task)
}

func traceHandler(w http.ResponseWriter, r *http.Request) {
	reqID := r.Header.Get("X-Request-ID")
	if reqID == "" {
		reqID = "req-unknown"
	}
	ctx := context.WithValue(r.Context(), requestIDKey, reqID)
	log.Printf("trace: handling request %s", requestIDFrom(ctx))
	writeJSON(w, http.StatusOK, map[string]string{"requestID": requestIDFrom(ctx)})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /ping", ping)
	mux.HandleFunc("GET /tasks", listTasks)
	mux.HandleFunc("POST /tasks", createTask)
	mux.HandleFunc("GET /tasks/{id}", getTask)
	mux.HandleFunc("GET /slow", slowHandler)
	mux.HandleFunc("GET /slow-store", slowStoreHandler)
	mux.HandleFunc("GET /trace", traceHandler)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
