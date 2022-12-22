package sqlite

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/samverrall/go-task-application/user-service/internal/port/domain"
	"github.com/samverrall/go-task-application/user-service/internal/port/domain/user"
	"github.com/samverrall/go-task-application/user-service/internal/port/repository"
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

type gormUser struct {
	gorm.Model
	UUID     string
	Email    string
	Password string
}

func (ur *UserRepo) Add(ctx context.Context, in user.User) (*user.User, error) {
	gormUser := userToGorm(in)
	if err := ur.db.Create(&gormUser).Error; err != nil {
		return nil, err
	}
	return &in, nil
}

func (ur *UserRepo) Get(ctx context.Context, uuid uuid.UUID) (*user.User, error) {
	result := gormUser{}
	err := ur.db.First(&result, "uuid = ?", uuid.String()).Error
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, repository.ErrUserNotFound

	case err != nil:
		return nil, err
	}
	return gormToUser(result), nil
}
