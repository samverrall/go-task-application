package user

import (
	"fmt"
	"unicode/utf8"
)

const (
	PasswordMinLength = 6
)

type Password string

func NewPassword(password string) (Password, error) {
	if utf8.RuneCountInString(password) < PasswordMinLength {
		return "", fmt.Errorf("a password must be greater than %d characters", PasswordMinLength)
	}
	return Password(password), nil
}

func (p Password) String() string {
	return string(p)
}
