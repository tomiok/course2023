package repository

import (
	"github.com/tomiok/course-23/db"
	"github.com/tomiok/course-23/internal/movie"
)

type Repository struct {
	db *db.MoviesDB
}

func (r Repository) CreateMovie(m movie.Movie) (movie.Movie, error) {
	return movie.Movie{}, nil
}
