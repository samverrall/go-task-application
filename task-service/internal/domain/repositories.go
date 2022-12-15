package domain

import (
	"context"

	"github.com/samverrall/task-service/internal/domain/task"
)

type TaskRepo interface {
	Add(ctx context.Context, t task.Task) (*task.Task, error)
}
