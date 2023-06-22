package repository

import (
	"github.com/rs/zerolog/log"
	"github.com/tomiok/course-23/db"
	"github.com/tomiok/course-23/internal/movie"
)

type Repository struct {
	db *db.MoviesDB
}

func NewRepo(db *db.MoviesDB) Repository {
	return Repository{
		db: db,
	}
}

func (r Repository) CreateMovie(m movie.Movie) (movie.Movie, error) {
	err := r.db.DB.Create(&m).Error

	if err != nil {
		return movie.Movie{}, err
	}

	log.Info().Msgf("movie %s saved OK to the DB", m.Title)
	return m, nil
}
