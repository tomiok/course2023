package repository

import (
	"fmt"
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
	fmt.Println(m)
	fmt.Println("Movie is saved")
	return movie.Movie{}, nil
}
