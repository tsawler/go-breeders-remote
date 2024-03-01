package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"net/http"
)

// routes sets up application routes.
func (app *application) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	mux.Get("/api/cat-breeds/all/json", app.GetAllCatBreedsJSON)
	mux.Get("/api/cat-breeds/all/xml", app.GetAllCatBreedsXML)

	mux.Get("/api/cat-breeds/{breed}/xml", app.GetCatBreedByNameXML)
	mux.Get("/api/cat-breeds/{breed}/json", app.GetCatBreedByNameJSON)

	return mux
}
