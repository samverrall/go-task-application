package mockrepo

import (
	"context"
	"errors"

	"github.com/samverrall/task-service/internal/port/domain"
	"github.com/samverrall/task-service/internal/port/domain/task"
)

type TaskRepoMock struct {
	tasks map[string]task.Task
}

func NewMockTaskRepo() domain.TaskRepo {
	return &TaskRepoMock{
		tasks: make(map[string]task.Task),
	}
}

func (t *TaskRepoMock) Add(ctx context.Context, task task.Task) (*task.Task, error) {
	if _, exists := t.tasks[task.UUID.String()]; exists {
		return nil, errors.New("mock: task already uuid exists")
	}

	t.tasks[task.UUID.String()] = task

	return nil, nil
}
