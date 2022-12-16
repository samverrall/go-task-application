package sqlite_test

import (
	"testing"

	"github.com/samverrall/task-service/internal/port/repository/repotest"
	"github.com/samverrall/task-service/internal/port/repository/sqlite"
	sqliteDB "github.com/samverrall/task-service/pkg/sqlite"
	"github.com/stretchr/testify/assert"
)

func TestTask(t *testing.T) {
	db, err := sqliteDB.Connect("file::memory:?cache=shared")
	assert.NoError(t, err, "new inmemory sqlite error")

	conn := db.GetDB()

	repo := sqlite.NewTaskRepo(conn)
	repotest.RunTaskTests(t, repo)
}
