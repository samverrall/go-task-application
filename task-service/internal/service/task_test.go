package service

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"testing/quick"

	"github.com/davecgh/go-spew/spew"
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

	execute := func(input *CreateTaskDTO) error {
		_, err := taskService.CreateTask(ctx, input)
		return err
	}

	t.Run("invalid name", func(t *testing.T) {
		f := func(in *CreateTaskDTO) bool {
			spew.Dump(in)
			err := execute(in)

			fmt.Println(errors.Is(err, domain.ErrInvalidTaskName))
			return errors.Is(err, domain.ErrInvalidTaskName)
		}
		quick.Check(f, &quick.Config{
			MaxCount: 100,
		})
	})
}
