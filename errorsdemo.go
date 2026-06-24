package main

import (
	"errors"
	"fmt"
)

// ValidationError is a custom error type: any type with an Error() string method
// satisfies the built-in error interface. A struct lets the error carry context
// (which field failed and why) that callers can read back out with errors.As.
type ValidationError struct {
	Field string
	Msg   string
}

// Error makes *ValidationError satisfy the error interface.
func (e *ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Msg)
}

// errorsDemo shows the Go error model: returning and checking errors, wrapping
// with %w, inspecting with errors.Is and errors.As, the store's sentinel error,
// and recovering from a panic.
func errorsDemo() {
	fmt.Println("\n-- Errors --")

	// loadTask wraps ErrTaskNotFound with context using %w. Printing the error
	// shows the whole chain; errors.Is still finds the sentinel underneath.
	if _, err := loadTask(99); err != nil {
		fmt.Println("loadTask(99) failed:", err)
		fmt.Println("  errors.Is(err, ErrTaskNotFound):", errors.Is(err, ErrTaskNotFound))
	}

	// validateTitle returns a *ValidationError. errors.As copies it back out so
	// we can read its Field and Msg, even if it were wrapped further.
	err := validateTitle("")
	var ve *ValidationError
	if errors.As(err, &ve) {
		fmt.Printf("validation failed on field %q: %s\n", ve.Field, ve.Msg)
	}

	// The store now returns ErrTaskNotFound directly for a missing id — this is
	// the lookup the HTTP layer maps to a 404 later in the course.
	store := NewInMemoryStore()
	if _, err := store.Find(1); errors.Is(err, ErrTaskNotFound) {
		fmt.Println("store.Find(1): task not found (sentinel matched)")
	}

	// safeDivide turns a divide-by-zero panic into a returned error, so a bug in
	// one place does not crash the whole program.
	if _, err := safeDivide(10, 0); err != nil {
		fmt.Println("safeDivide(10, 0) recovered:", err)
	}
}

// loadTask pretends to read a task from somewhere that can fail, then wraps the
// failure with context using %w so callers can still unwrap it with errors.Is.
func loadTask(id int) (Task, error) {
	store := NewInMemoryStore()
	t, err := store.Find(id)
	if err != nil {
		return Task{}, fmt.Errorf("load task %d: %w", id, err)
	}
	return t, nil
}

// validateTitle returns a *ValidationError when the title is empty.
func validateTitle(title string) error {
	if title == "" {
		return &ValidationError{Field: "title", Msg: "must not be empty"}
	}
	return nil
}

// safeDivide recovers from a divide-by-zero panic and returns it as an error, so
// a bug in one place does not crash the whole program. The named return err is
// set inside the deferred recover.
func safeDivide(a, b int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recovered from panic: %v", r)
		}
	}()
	return a / b, nil
}
