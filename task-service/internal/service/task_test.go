package service

import (
	"context"
	"testing"
	"testing/quick"

	"github.com/samverrall/go-task-application/logger"
	"github.com/samverrall/task-service/internal/domain"
)

type mockTaskRepo struct{}

func (mt mockTaskRepo) CreateTask(ctx context.Context, t *domain.Task) (*domain.Task, error) {
	return t, nil
}

func TestCreateServive(t *testing.T) {
	ctx := context.Background()
	l := logger.New("debug")
	taskService := NewTaskService(mockTaskRepo{}, l)

	execute := func(name domain.TaskName, completeBy domain.TaskCompleteBy) error {
		_, err := taskService.CreateTask(ctx, &CreateTaskDTO{
			Name:       name.String(),
			CompleteBy: completeBy.Time(),
		})
		return err
	}

	t.Run("valid inputs", func(t *testing.T) {
		f := func(name domain.TaskName, completeBy domain.TaskCompleteBy) bool {
			err := execute(name, completeBy)
			return err == nil
		}
		quick.Check(f, nil)
	})
}
