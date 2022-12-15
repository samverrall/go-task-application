package task

import (
	"context"
	"time"

	"github.com/samverrall/task-service/internal/domain"
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

func (ts *TaskService) CreateTask(ctx context.Context, taskDTO CreateTaskDTO) (*domain.Task, error) {
	ts.logger.Info("ts.CreateTask Invoked")

	taskName, err := domain.NewTaskName(taskDTO.Name)
	if err != nil {
		return nil, err
	}

	taskCompleteBy, err := domain.NewTaskCompleteBy(taskDTO.CompleteBy)
	if err != nil {
		return nil, err
	}

	task := domain.NewTask(taskName, taskCompleteBy)

	return ts.repo.Add(ctx, task)
}
