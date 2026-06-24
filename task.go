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
// so it uses a value receiver (t Task): it only reads the fields and never
// changes them.
func (t Task) Summary() string {
	status := "open"
	if t.Done {
		status = "done"
	}
	return fmt.Sprintf("#%d %q [%s]", t.ID, t.Title, status)
}

// markDoneByValue tries to mark the task done through a value receiver. The
// receiver is a copy, so flipping Done here changes only the copy and the
// caller's task is untouched. This is the classic beginner trap. It returns the
// copy's Done so the change is visible locally even though it never sticks.
func (t Task) markDoneByValue() bool {
	t.Done = true
	return t.Done
}

// MarkDone marks the task done through a pointer receiver (t *Task). Because the
// receiver is the address of the original task, the change sticks.
func (t *Task) MarkDone() {
	t.Done = true
}
