package main

import (
	"github.com/tomiok/course-23/db"
	"github.com/tomiok/course-23/internal/movie"
	"github.com/tomiok/course-23/internal/movie/handler"
	"github.com/tomiok/course-23/internal/movie/repository"
)

type Dependencies struct {
	movieHandler *handler.Handler
}

func newDeps() *Dependencies {
	moviesDB, err := db.CreateDB()

	if err != nil {
		panic(err)
	}

	movieStorage := repository.NewRepo(moviesDB)
	movieService := movie.NewService(movieStorage)
	movieHandler := handler.NewHandler(movieService)

	return &Dependencies{
		movieHandler: movieHandler,
	}
}
