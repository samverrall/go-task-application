package sqlite

import (
	"context"

	"github.com/google/uuid"
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

type gormUser struct {
	gorm.Model
	Email string
}

func (ur *UserRepo) Get(ctx context.Context, uuid uuid.UUID) (*user.User, error) {
	result := gormUser{}
	if err := ur.db.Where("uuid = ?", uuid.String()).Find(&result).Error; err != nil {
		return nil, err
	}
	return gormToUser(result), nil
}
