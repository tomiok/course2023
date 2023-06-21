// Package db package db holds all the information regarding DBs
package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func CreateDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		// built-in function panic()
		return nil, err
	}

	return db, nil
}
