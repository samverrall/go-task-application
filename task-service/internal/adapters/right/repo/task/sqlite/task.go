package sqlite

import (
	"context"
	"time"

	"github.com/samverrall/task-service/internal/port/domain"
	"github.com/samverrall/task-service/internal/port/domain/task"
	"gorm.io/gorm"
)

type TaskRepo struct {
	db *gorm.DB
}

func NewTaskRepo(db *gorm.DB) domain.TaskRepo {
	return &TaskRepo{
		db: db,
	}
}

type GormTask struct {
	gorm.Model
	UUID       string
	Name       string
	CreatedAt  time.Time
	CompleteBy time.Time
}

func domainToGORM(t task.Task) *GormTask {
	return &GormTask{
		UUID:       t.UUID.String(),
		Name:       t.Name.String(),
		CreatedAt:  t.CreatedAt,
		CompleteBy: t.CompleteBy.Time(),
	}
}

func (tr *TaskRepo) Add(ctx context.Context, t task.Task) (*task.Task, error) {
	if err := tr.db.Create(domainToGORM(t)).Error; err != nil {
		return nil, err
	}
	return &t, nil
}
