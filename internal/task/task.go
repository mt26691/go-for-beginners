package task

import (
	"errors"
	"strings"
	"time"
)

// MaxTitleLength caps how long a task title may be.
const MaxTitleLength = 200

// Validation errors returned by Task.Validate. They are sentinels so callers
// could match them with errors.Is, but the handler just uses their message.
var (
	ErrTitleRequired = errors.New("title is required")
	ErrTitleTooLong  = errors.New("title must be 200 characters or fewer")
)

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	Done        bool      `json:"done"`
	CreatedAt   time.Time `json:"createdAt"`
}

// Validate trims whitespace from the title and description, then checks the
// business rules: a title is required and may not exceed MaxTitleLength. It
// mutates the receiver so the caller stores the trimmed values. The checks are
// written by hand here so the mechanics are visible; a larger project might use
// a library such as github.com/go-playground/validator instead.
func (t *Task) Validate() error {
	t.Title = strings.TrimSpace(t.Title)
	t.Description = strings.TrimSpace(t.Description)

	if t.Title == "" {
		return ErrTitleRequired
	}
	if len(t.Title) > MaxTitleLength {
		return ErrTitleTooLong
	}
	return nil
}
