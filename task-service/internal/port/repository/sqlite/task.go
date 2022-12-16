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

func NewTaskRepo(db *gorm.DB) (domain.TaskRepo, error) {
	if err := migrate(db); err != nil {
		return nil, err
	}
	return &TaskRepo{
		db: db,
	}, nil
}

func migrate(db *gorm.DB) error {
	return db.AutoMigrate(gormTask{})
}

type gormTask struct {
	gorm.Model
	UUID       string
	Name       string
	CreatedAt  time.Time
	CompleteBy time.Time
}

func domainToGORM(t task.Task) *gormTask {
	return &gormTask{
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
