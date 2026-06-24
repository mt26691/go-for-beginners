package main

import "fmt"

// Task is the running example for the rest of the course: a single to-do item
// with an ID, a Title, and whether it is Done. The API chapters reuse this exact
// struct.
type Task struct {
	ID    int
	Title string
	Done  bool
}

// Summary returns a one-line description of the task. It is a read-only method,
// so it uses a value receiver (t Task).
func (t Task) Summary() string {
	// TODO: return something like: #1 "write code" [open]
	return fmt.Sprintf("#%d %q [todo]", t.ID, t.Title)
}

// markDoneByValue tries to mark the task done through a value receiver. The
// receiver is a copy, so the change will not stick on the caller's task.
func (t Task) markDoneByValue() bool {
	// TODO: set t.Done = true and return it, then see why the caller is untouched
	return t.Done
}

// MarkDone marks the task done through a pointer receiver (t *Task) so the
// change sticks.
func (t *Task) MarkDone() {
	// TODO: set t.Done = true
}
