package sqlite

import (
	"context"

	"github.com/samverrall/go-task-application/user-service/internal/port/domain"
	"github.com/samverrall/go-task-application/user-service/internal/port/domain/user"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) (domain.UserRepo, error) {
	if err := migrate(db); err != nil {
		return nil, err
	}
	return &UserRepo{
		db: db,
	}, nil
}

func migrate(db *gorm.DB) error {
	return db.AutoMigrate(gormUser{})
}

type gormUser struct{}

func (ur *UserRepo) CreateUser(ctx context.Context, u user.User) error {
	return nil
}
