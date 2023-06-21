package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.With(SayHelloMiddleware).Get("/movies", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("GET, welcome"))
	})
	r.Post("/movies", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("POST, welcome"))
	})
	r.Delete("/movies", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("DELETE, welcome"))
	})
	http.ListenAndServe(":8080", r)
}

func SayHelloMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("hello there: " + r.Method)
		next.ServeHTTP(w, r)
	})
}

// rewrite the last exercise using CHI.
// use only one URL (in plural) with the three verbs (GET, POST, DELETE)

// addMovieHandler
// readMovieHandler
