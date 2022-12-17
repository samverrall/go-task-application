package user

import (
	"fmt"
)

type Password string

func NewPassword(password string) (Password, error) {
	if len(password) < PasswordMinLength {
		return "", fmt.Errorf("a password must be greater than %d characters", PasswordMinLength)
	}
	return Password(password), nil
}

func (p Password) String() string {
	return string(p)
}
