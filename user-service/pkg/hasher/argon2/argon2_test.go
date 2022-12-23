package argon2_test

import (
	"testing"

	"github.com/samverrall/go-task-application/user-service/pkg/hasher/argon2"
)

func TestCompare(t *testing.T) {
	a := argon2.New()

	password := "password"
	t.Run("password successfully matches", func(t *testing.T) {
		passwordHashed, err := a.Generate(password)
		if err != nil {
			t.Error(err)
		}

		match, err := a.Compare(password, passwordHashed)
		if err != nil {
			t.Error(err)
		}

		if !match {
			t.Error("password compare does not match")
		}
	})

	t.Run("password mismatch", func(t *testing.T) {
		passwordHashed, err := a.Generate(password)
		if err != nil {
			t.Error(err)
		}

		match, err := a.Compare("different", passwordHashed)
		if err != nil {
			t.Error(err)
		}

		if match {
			t.Error("password compare match")
		}
	})
}
