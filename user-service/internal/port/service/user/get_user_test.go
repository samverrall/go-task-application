package user_test

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/samverrall/go-task-application/logger"
	userDomain "github.com/samverrall/go-task-application/user-service/internal/port/domain/user"
	"github.com/samverrall/go-task-application/user-service/internal/port/repository"
	"github.com/samverrall/go-task-application/user-service/internal/port/repository/repotest"
	"github.com/samverrall/go-task-application/user-service/internal/port/repository/sqlite"
	"github.com/samverrall/go-task-application/user-service/internal/port/service/user"
	"github.com/samverrall/go-task-application/user-service/pkg/hasher/argon2"
	sqliteconn "github.com/samverrall/go-task-application/user-service/pkg/sqlite"
	"github.com/stretchr/testify/assert"
)

type canGetUserGuard struct {
	canGetUserGuard bool
}

func (g canGetUserGuard) CanGetUser() bool {
	return g.canGetUserGuard
}

func TestGetUser(t *testing.T) {
	ctx := context.Background()

	canGetUser := canGetUserGuard{canGetUserGuard: true}
	cannotGetUser := canGetUserGuard{canGetUserGuard: false}

	db, err := sqliteconn.Connect("file::memory:?cache=shared")
	assert.NoError(t, err, "failed to create sqlite inmemory connection: %v", err)

	repo, err := sqlite.NewUserRepo(db.GetDB())
	assert.NoError(t, err, "failed to create sqlite inmemory connection: %v", err)

	userSvc := user.NewService(repo, logger.New("debug"), argon2.New())

	existingUser, err := repotest.AddUser(ctx, t, repo, "test@test.com", "test123")
	assert.NoError(t, err, "failed to a user: %v", err)

	t.Run("successfully gets user", func(t *testing.T) {
		user, err := userSvc.GetUser(ctx, user.GetUserRequest{
			UserUUID: existingUser.UUID.String(),
		}, canGetUser)
		if err != nil {
			t.Error(err)
		}
		if existingUser.Email != userDomain.Email(user.Email) {
			t.Errorf("wanted user email %s, got %s", existingUser.Email.String(), user.Email)
		}
		if existingUser.UUID != uuid.MustParse(user.UUID) {
			t.Errorf("wanted user uuid %s, got %s", existingUser.Email.String(), user.Email)
		}
	})

	t.Run("unknown user returns not found", func(t *testing.T) {
		_, err := userSvc.GetUser(ctx, user.GetUserRequest{
			UserUUID: uuid.NewString(),
		}, canGetUser)
		if !errors.Is(err, repository.ErrUserNotFound) {
			t.Errorf("want %v, got %v", repository.ErrUserNotFound, err)
		}
	})

	t.Run("unauthorised user returns an error", func(t *testing.T) {
		_, err := userSvc.GetUser(ctx, user.GetUserRequest{
			UserUUID: uuid.NewString(),
		}, cannotGetUser)
		if err == nil {
			t.Error("want error, got <nil>")
		}
	})

	t.Run("invalid uuid input failure", func(t *testing.T) {
		_, err := userSvc.GetUser(ctx, user.GetUserRequest{
			UserUUID: "invalid-uuid",
		}, canGetUser)
		if err == nil {
			t.Error("want err, got nil", err)
		}
	})
}
