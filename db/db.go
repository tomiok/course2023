// Package db package db holds all the information regarding DBs
package db

import (
	"github.com/rs/zerolog/log"
	"github.com/tomiok/course-23/internal/movie"
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

	migrate(db)
	return &MoviesDB{DB: db}, nil
}

func migrate(db *gorm.DB) {
	log.Error().Msgf("%s", db.AutoMigrate(&movie.Movie{}))
	log.Error().Msgf("%s", db.AutoMigrate(&movie.Actor{}))
}
