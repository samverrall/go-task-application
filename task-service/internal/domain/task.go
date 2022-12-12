package domain

import (
	"errors"
	"strings"
	"time"

	"gorm.io/gorm"
)

var (
	ErrInvalidTaskName = errors.New("a task requires a name")
)

type Task struct {
	gorm.Model
	UUID       string    `json:"uuid"`
	Name       string    `json:"name"`
	CreatedAt  time.Time `json:"createdAt"`
	CompleteBy time.Time `json:"completeBy"`
}

func (t *Task) Validate() error {
	if strings.TrimSpace(t.Name) == "" {
		return ErrInvalidTaskName
	}
	return nil
}
