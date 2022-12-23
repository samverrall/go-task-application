package user

import (
	"fmt"
	"unicode/utf8"

	"github.com/samverrall/go-task-application/user-service/pkg/hasher"
)

const (
	PasswordMinLength = 6
)

type HashedPassword string

func NewHashedPassword(password string, hasher hasher.Hasher) (HashedPassword, error) {
	if utf8.RuneCountInString(password) < PasswordMinLength {
		return "", fmt.Errorf("a password must be greater than %d characters", PasswordMinLength)
	}

	hashed, err := hasher.Generate(password)
	if err != nil {
		return "", err
	}

	return HashedPassword(hashed), nil
}

func (p HashedPassword) String() string {
	return string(p)
}
