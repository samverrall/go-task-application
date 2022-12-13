package sqlite

import (
	"fmt"

	"github.com/samverrall/go-task-application/user-service/internal/domain"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SQLiteDB struct {
	db *gorm.DB
}

func Connect(databaseDir string) (*SQLiteDB, error) {
	db, err := gorm.Open(sqlite.Open(databaseDir), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("%w: failed to open gorm connection", err)
	}
	return &SQLiteDB{
		db: db,
	}, nil
}

func (sq *SQLiteDB) GetDB() *gorm.DB {
	return sq.db
}

func (sq *SQLiteDB) Migrate() error {
	return sq.db.AutoMigrate(domain.User{})
}
