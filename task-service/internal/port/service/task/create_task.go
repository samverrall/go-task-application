package task

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/samverrall/task-service/internal/port/domain/task"
)

type CreateTaskGuard interface {
	CanCreateTask() bool
}

// CreateTaskDTO  is a middleman 'DTO' (Data Transfer Object) to decouple
// the domains from our ports. This way the port can adapt to the inputs of adapters.
type CreateTaskDTO struct {
	Name       string
	CompleteBy time.Time
}

func (ts *TaskService) CreateTask(ctx context.Context, taskDTO CreateTaskDTO) (*task.Task, error) {
	ts.logger.Info("ts.CreateTask Invoked")

	taskName, err := task.NewName(taskDTO.Name)
	if err != nil {
		return nil, err
	}

	taskCompleteBy, err := task.NewCompleteBy(taskDTO.CompleteBy)
	if err != nil {
		return nil, err
	}

	id := uuid.New()
	task := task.New(id, taskName, taskCompleteBy)

	return ts.repo.Add(ctx, task)
}
