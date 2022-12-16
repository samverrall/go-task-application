package repotest

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/samverrall/task-service/internal/port/domain"
	"github.com/samverrall/task-service/internal/port/domain/task"
	"github.com/stretchr/testify/assert"
)

func RunTaskTests(t *testing.T, tasks domain.TaskRepo) {
	ctx := context.Background()
	tomorrow := time.Now().AddDate(0, 0, 1)

	_, err := addTask(t, ctx, tasks, "foo", tomorrow)
	assert.NoError(t, err, "addTask failure")
}

func addTask(t *testing.T, ctx context.Context, tasks domain.TaskRepo, name string, completeBy time.Time) (*task.Task, error) {
	t.Helper()

	taskName, err := task.NewName(name)
	if err != nil {
		t.Error(err)
	}

	taskCompleteBy, err := task.NewCompleteBy(completeBy)
	if err != nil {
		t.Error(err)
	}

	newTask := task.New(uuid.New(), taskName, taskCompleteBy)

	return tasks.Add(ctx, newTask)
}
