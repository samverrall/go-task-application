package service

import (
	"context"

	"github.com/samverrall/go-task-application/logger"
	"github.com/samverrall/task-service/internal/domain"
)

type TaskServicer interface {
	CreateTask(ctx context.Context, taskDTO *CreateTaskDTO) (*domain.Task, error)
}

type TaskService struct {
	repo   domain.TaskRepo
	logger logger.Logger
}

// CreateTaskDTO is a middleman 'DTO' (Data Transfer Object) to decouple
// the domains from our ports. This way the port can adapt to the inputs of adapters.
type CreateTaskDTO struct {
	Name string
}

func NewTaskService(repo domain.TaskRepo, log logger.Logger) TaskServicer {
	return &TaskService{
		repo:   repo,
		logger: log,
	}
}

func (ts *TaskService) CreateTask(ctx context.Context, taskDTO *CreateTaskDTO) (*domain.Task, error) {
	ts.logger.Info("ts.CreateTask Invoked")

	taskName, err := domain.NewTaskName(taskDTO.Name)
	if err != nil {
		return nil, err
	}

	task := domain.NewTask(taskName)

	return ts.repo.CreateTask(ctx, task)
}
