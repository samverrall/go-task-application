package task

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

const (
	taskNameMinLength = 3
)

var (
	ErrInvalidTaskName  = errors.New("empty task name is not allowed")
	ErrTaskNameTooSmall = fmt.Errorf("task name must be greater than %d", taskNameMinLength)

	ErrCompleteByTaskEmpty = errors.New("complete by date cannot be empty")
	ErrCompleteByInPast    = errors.New("task complete by must be in the future")
)

type Task struct {
	UUID       uuid.UUID
	Name       Name
	CreatedAt  time.Time
	CompleteBy CompleteBy
}

type Name string

type CompleteBy time.Time

func New(name Name, completeBy CompleteBy) *Task {
	return &Task{
		Name:       name,
		CompleteBy: completeBy,
	}
}

// NewTaskName creates and handles validation for new task name.
// If an error is returned an invalid name has been supplied.
// This means you don't have to copy name validation all over the place, and you
// must construct the TaskName type to pass into a NewTask.
// Which encforces some validation before the aggregate root is even created.
// It also means the aggregate root is free to focus only on business logic,
// which data validation is not.
// All the input validation is centralised in one place.
func NewName(name string) (Name, error) {
	if strings.TrimSpace(name) == "" {
		return "", ErrInvalidTaskName
	}

	if len(name) < taskNameMinLength {
		return "", ErrTaskNameTooSmall
	}

	return Name(name), nil
}

func (tn Name) String() string {
	return string(tn)
}

func NewCompleteBy(completeBy time.Time) (CompleteBy, error) {
	out := CompleteBy{}

	if completeBy.IsZero() {
		return out, ErrCompleteByTaskEmpty
	}

	if completeBy.Before(time.Now()) {
		return out, ErrCompleteByInPast
	}

	return CompleteBy(completeBy), nil
}

func (cb CompleteBy) Time() time.Time {
	return time.Time(cb)
}
