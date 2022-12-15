package sqlite_test

import (
	"testing"

	"github.com/samverrall/task-service/internal/repository/sqlite"
	"github.com/samverrall/task-service/internal/repository/sqlite/repotest"
	"github.com/stretchr/testify/assert"
	driver "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestTask(t *testing.T) {
	db, err := gorm.Open(driver.Open("file::memory:?cache=shared"), nil)
	assert.NoError(t, err, "new inmemory sqlite error")

	err = db.AutoMigrate(sqlite.GormTask{})
	assert.NoError(t, err, "auto migrate err")

	repo := sqlite.NewTaskRepo(db)
	repotest.RunTaskTests(t, repo)
}
