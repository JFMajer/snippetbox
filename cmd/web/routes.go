package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not to be exceeded
	})

	mux.Use(cors.Handler)
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Logger)
	mux.Use(secureHeaders)
	mux.Use(middleware.Heartbeat("/ping"))

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	mux.Get("/", app.home)
	mux.Get("/snippet/view/{id}", app.snippetView)
	mux.Post("/snippet/create", app.snippetCreatePost)
	mux.Get("/snippet/create", app.snippetCreateForm)

	return mux
}
