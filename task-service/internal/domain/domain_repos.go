package domain

import "context"

type TaskRepo interface {
	Add(ctx context.Context, t *Task) (*Task, error)
}
