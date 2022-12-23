package user_test

import (
	"context"
	"testing"

	"github.com/samverrall/go-task-application/logger"
	userDomain "github.com/samverrall/go-task-application/user-service/internal/port/domain/user"
	"github.com/samverrall/go-task-application/user-service/internal/port/repository/sqlite"
	"github.com/samverrall/go-task-application/user-service/internal/port/service/user"
	"github.com/samverrall/go-task-application/user-service/pkg/hasher/argon2"
	sqliteconn "github.com/samverrall/go-task-application/user-service/pkg/sqlite"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	ctx := context.Background()

	db, err := sqliteconn.Connect("file::memory:?cache=shared")
	assert.NoError(t, err, "failed to create sqlite inmemory connection: %v", err)

	repo, err := sqlite.NewUserRepo(db.GetDB())
	assert.NoError(t, err, "failed to create sqlite inmemory connection: %v", err)

	userSvc := user.NewService(repo, logger.New("debug"), argon2.New())

	t.Run("successfully registers user", func(t *testing.T) {
		email := "test@test.com"
		err := userSvc.Register(ctx, user.RegisterRequest{
			Email:    "test@test.com",
			Password: "password",
		})
		if err != nil {
			t.Error(err)
		}

		_, err = repo.GetByEmail(ctx, userDomain.Email(email))
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("invalid password input failure", func(t *testing.T) {
		invalidPassword := "12"
		err := userSvc.Register(ctx, user.RegisterRequest{
			Email:    "test@test.com",
			Password: invalidPassword,
		})
		if err == nil {
			t.Error("want err, got nil", err)
		}
	})

	t.Run("invalid email input failure", func(t *testing.T) {
		invalidEmail := "test2.com"
		err := userSvc.Register(ctx, user.RegisterRequest{
			Email:    invalidEmail,
			Password: "password",
		})
		if err == nil {
			t.Error("want err, got nil", err)
		}
	})
}
