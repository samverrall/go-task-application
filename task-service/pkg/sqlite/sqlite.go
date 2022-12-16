package sqlite

import (
	"fmt"

	driver "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SQLiteDB struct {
	db *gorm.DB
}

func Connect(database string) (*SQLiteDB, error) {
	db, err := gorm.Open(driver.Open(database), &gorm.Config{})
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
	return nil
}
