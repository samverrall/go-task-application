package repotest

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/samverrall/go-task-application/user-service/internal/port/domain"
	"github.com/samverrall/go-task-application/user-service/internal/port/domain/user"
	"github.com/samverrall/go-task-application/user-service/pkg/hasher/argon2"
	"github.com/stretchr/testify/assert"
)

func RunUserTests(t *testing.T, tasks domain.UserRepo) {
	ctx := context.Background()

	newUser, err := addUser(ctx, t, tasks, "foo@foo.com", "password")
	assert.NoError(t, err, "addUser failure")

	_, err = getUser(ctx, t, tasks, newUser.UUID.String())
	assert.NoError(t, err, "getUser failure")
}

func getUser(ctx context.Context, t *testing.T, users domain.UserRepo, uuidIn string) (*user.User, error) {
	t.Helper()

	userUuid, err := uuid.Parse(uuidIn)
	if err != nil {
		return nil, err
	}

	return users.Get(ctx, userUuid)
}

func addUser(ctx context.Context, t *testing.T, users domain.UserRepo, email, password string) (*user.User, error) {
	t.Helper()

	userEmail, err := user.NewEmail(email)
	if err != nil {
		t.Error(err)
		return nil, err
	}

	hasher := argon2.New()
	userPassword, err := user.NewHashedPassword(password, hasher)
	if err != nil {
		t.Error(err)
		return nil, err
	}

	newUser := user.New(uuid.New(), userEmail, userPassword)

	return users.Add(ctx, newUser)
}
