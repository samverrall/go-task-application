package sqlite_test

import (
	"testing"

	"github.com/samverrall/go-task-application/user-service/internal/adapters/right/repo/user/repotest"
	"github.com/samverrall/go-task-application/user-service/internal/adapters/right/repo/user/sqlite"
	sqliteDB "github.com/samverrall/task-service/pkg/sqlite"
	"github.com/stretchr/testify/assert"
)

func TestTask(t *testing.T) {
	db, err := sqliteDB.Connect("file::memory:?cache=shared")
	assert.NoError(t, err, "new inmemory sqlite error")

	conn := db.GetDB()

	repo, err := sqlite.NewUserRepo(conn)
	assert.NoError(t, err, "failed to make new task repo")

	repotest.RunUserTests(t, repo)
}
