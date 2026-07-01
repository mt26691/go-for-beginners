// main.go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	Done        bool   `json:"done"`
	Internal    string `json:"-"`
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Printf("writeJSON: %v", err)
	}
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
	task := Task{ID: id, Title: "Sample task", Done: false}
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

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /ping", ping)
	mux.HandleFunc("GET /tasks", listTasks)
	mux.HandleFunc("POST /tasks", createTask)
	mux.HandleFunc("GET /tasks/{id}", getTask)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
