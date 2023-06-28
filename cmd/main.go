package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
)

const port = ":8080"

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)

	deps := newDeps()
	r.Post("/movies", deps.movieHandler.CreateMovieHandler)

	log.Info().Msgf("web app is running in port %s", port)
	http.ListenAndServe(port, r)
}
