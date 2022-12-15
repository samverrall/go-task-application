package task

import (
	"context"
	"testing"
	"testing/quick"

	"github.com/samverrall/go-task-application/logger"
	"github.com/samverrall/task-service/internal/domain/task"
)

type mockTaskRepo struct{}

func (mt mockTaskRepo) Add(ctx context.Context, t *task.Task) (*task.Task, error) {
	return t, nil
}

func TestCreateServive(t *testing.T) {
	ctx := context.Background()
	l := logger.New("debug")
	taskService := NewService(mockTaskRepo{}, l)

	execute := func(name task.Name, completeBy task.CompleteBy) error {
		_, err := taskService.CreateTask(ctx, CreateTaskDTO{
			Name:       name.String(),
			CompleteBy: completeBy.Time(),
		})
		return err
	}

	t.Run("valid inputs", func(t *testing.T) {
		f := func(name task.Name, completeBy task.CompleteBy) bool {
			err := execute(name, completeBy)
			return err == nil
		}
		quick.Check(f, nil)
	})
}
