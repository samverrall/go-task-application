package repository

import (
	"context"

	"github.com/samverrall/task-service/internal/domain"
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

func (tr *TaskRepo) CreateTask(ctx context.Context, t *domain.Task) (*domain.Task, error) {
	if err := tr.db.Create(t).Error; err != nil {
		return nil, err
	}
	return t, nil
}
