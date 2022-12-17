package sqlite

import (
	"github.com/google/uuid"
	"github.com/samverrall/go-task-application/user-service/internal/port/domain/user"
)

func gormToUser(u gormUser) *user.User {
	return &user.User{
		UUID:  uuid.MustParse(u.UUID),
		Email: user.Email(u.Email),
	}
}

func userToGorm(u user.User) gormUser {
	return gormUser{
		UUID:     u.UUID.String(),
		Email:    u.Email.String(),
		Password: u.Password.String(),
	}
}
