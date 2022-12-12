package api

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/samverrall/task-service/internal/domain"
	"github.com/samverrall/task-service/pkg/logger"
)

type TaskServicer interface {
	CreateTask(ctx context.Context, t *domain.Task) (*domain.Task, error)
}

type TaskService struct {
	repo   domain.TaskRepo
	logger logger.Logger
}

func NewTaskService(repo domain.TaskRepo, log logger.Logger) TaskServicer {
	return &TaskService{
		repo:   repo,
		logger: log,
	}
}

func (ts *TaskService) CreateTask(ctx context.Context, t *domain.Task) (*domain.Task, error) {
	ts.logger.Info("ts.CreateTask Invoked")

	if err := t.Validate(); err != nil {
		ts.logger.Error("Invalid task supplied: %s", err.Error())
		return nil, err
	}

	t.UUID = uuid.NewString()

	if t.CreatedAt.IsZero() {
		t.CreatedAt = time.Now()
	}

	return ts.repo.CreateTask(ctx, t)
}