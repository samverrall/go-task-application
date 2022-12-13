package repository

import (
	"context"

	"github.com/samverrall/go-task-application/user-service/internal/domain"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) domain.UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (ur *UserRepo) CreateUser(ctx context.Context, u *domain.User) error {
	return nil
}
