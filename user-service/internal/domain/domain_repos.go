package domain

import "context"

type UserRepo interface {
	CreateUser(context.Context, *User) error
}
