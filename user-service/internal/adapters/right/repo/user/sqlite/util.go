package sqlite

import "github.com/samverrall/go-task-application/user-service/internal/port/domain/user"

func gormToUser(u gormUser) *user.User {
	return &user.User{
		Email: user.Email(u.Email),
	}
}
