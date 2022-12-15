package task

import (
	"context"

	"github.com/samverrall/go-task-application/logger"
	"github.com/samverrall/task-service/internal/domain"
	"github.com/samverrall/task-service/internal/domain/task"
)

type API interface {
	CreateTask(ctx context.Context, taskDTO CreateTaskDTO) (*task.Task, error)
}

type TaskService struct {
	repo   domain.TaskRepo
	logger logger.Logger
}

func NewService(repo domain.TaskRepo, log logger.Logger) API {
	return &TaskService{
		repo:   repo,
		logger: log,
	}
}
