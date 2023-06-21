// Package db package db holds all the information regarding DBs
package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type MoviesDB struct {
	DB *gorm.DB
}

func CreateDB() (*MoviesDB, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		// built-in function panic()
		return nil, err
	}

	return &MoviesDB{DB: db}, nil
}
