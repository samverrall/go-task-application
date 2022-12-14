package domain

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
)

type Task struct {
	UUID       uuid.UUID
	Name       TaskName
	CreatedAt  time.Time
	CompleteBy time.Time
}

type TaskName string

func NewTask(name TaskName) *Task {
	return &Task{
		Name: name,
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
func NewTaskName(name string) (TaskName, error) {
	if strings.TrimSpace(name) == "" {
		return "", ErrInvalidTaskName
	}

	if len(name) < taskNameMinLength {
		return "", ErrTaskNameTooSmall
	}

	return TaskName(name), nil
}

func (tn TaskName) String() string {
	return string(tn)
}
