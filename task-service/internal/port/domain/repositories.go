package domain

import (
	"context"

	"github.com/samverrall/task-service/internal/port/domain/task"
)

type TaskRepo interface {
	Add(ctx context.Context, t task.Task) (*task.Task, error)
}
