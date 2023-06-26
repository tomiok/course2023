package handler

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"github.com/tomiok/course-23/internal/movie"
	"net/http"
)

type Handler struct {
	*movie.Service
}

func NewHandler(s *movie.Service) *Handler {
	return &Handler{
		Service: s,
	}
}

func (h *Handler) CreateMovieHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// 1st I need to get the body.
	body := r.Body
	// 2nd deferred close of the body.
	defer body.Close()
	// 3rd read as json
	var _movie movie.Movie
	err := json.NewDecoder(body).Decode(&_movie)

	// err handling
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Error().Msgf("cannot decode request, %s", err)
		_ = json.NewEncoder(w).Encode(struct {
			Err string `json:"error"`
		}{Err: "cannot decode request"})
		return
	}

	// 4th use the json in the service.
	movieCreated, err := h.Create(_movie)

	// err handling
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Error().Msgf("cannot save movie, %s", err)
		_ = json.NewEncoder(w).Encode(struct {
			Err string `json:"err"`
		}{Err: "cannot save movies"})
		return
	}

	// finally
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(movieCreated)
}
