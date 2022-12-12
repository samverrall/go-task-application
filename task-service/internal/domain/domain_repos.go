package domain

import "context"

type TaskRepo interface {
	CreateTask(ctx context.Context, t *Task) (*Task, error)
}
