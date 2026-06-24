package main

import (
	"fmt"
)

// ValidationError is a custom error type: any type with an Error() string method
// satisfies the built-in error interface. A struct lets the error carry context
// (which field failed and why) that callers can read back out with errors.As.
type ValidationError struct {
	Field string
	Msg   string
}

// Error makes ValidationError satisfy the error interface.
// TODO: format Field and Msg into a readable message.
func (e *ValidationError) Error() string {
	return ""
}

// errorsDemo shows the Go error model: returning and checking errors, wrapping
// with %w, inspecting with errors.Is and errors.As, the store's sentinel error,
// and recovering from a panic.
func errorsDemo() {
	fmt.Println("\n-- Errors --")

	// TODO: call loadTask(99), print the wrapped error, and show errors.Is
	// finding ErrTaskNotFound through the wrap.
	_, err := loadTask(99)
	fmt.Println("loadTask(99):", err)

	// TODO: validate an empty title and use errors.As to pull the
	// *ValidationError back out for its Field and Msg.
	err = validateTitle("")
	fmt.Println("validateTitle(\"\"):", err)

	// TODO: look a missing task up in the store and show
	// errors.Is(err, ErrTaskNotFound).
	store := NewInMemoryStore()
	_, err = store.Find(1)
	fmt.Println("store.Find(1):", err)

	// TODO: call safeDivide(10, 0) to show recover turning a panic into an error.
	_, err = safeDivide(10, 0)
	fmt.Println("safeDivide(10, 0):", err)
}

// loadTask pretends to read a task from somewhere that can fail, then wraps the
// failure with context using %w so callers can still unwrap it.
// TODO: wrap ErrTaskNotFound with fmt.Errorf and %w.
func loadTask(id int) (Task, error) {
	return Task{}, nil
}

// validateTitle returns a *ValidationError when the title is empty.
// TODO: return a *ValidationError for an empty title, nil otherwise.
func validateTitle(title string) error {
	return nil
}

// safeDivide recovers from a divide-by-zero panic and returns it as an error, so
// a bug in one place does not crash the whole program.
// TODO: actually divide a by b inside a deferred recover so b == 0 returns an
// error instead of crashing. For now it returns zero values so make run stays green.
func safeDivide(a, b int) (result int, err error) {
	_ = a
	_ = b
	return 0, nil
}
