package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/tsawler/toolbox"
	"go-breeders-remote/models"
	"net/http"
)

// GetAllCatBreedsJSON gets a list of all cat breeds from the database and returns it as JSON.
func (app *application) GetAllCatBreedsJSON(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools

	// Get all cat breeds from the database.
	catBreeds, err := app.App.Models.CatBreed.All()
	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	// Write it out as JSON.
	_ = t.WriteJSON(w, http.StatusOK, catBreeds)
}

// GetAllCatBreedsXML gets a list of all cat breeds from the database and returns it as XML.
func (app *application) GetAllCatBreedsXML(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools

	// Get all breeds from the database.
	allBreeds, err := app.App.Models.CatBreed.All()
	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	// Since we are sending a slice, we need a wrapper, or we will not have a root element.
	type catBreeds struct {
		XMLName struct{}           `xml:"cat-breeds"` // this sets the name of the root element
		Breeds  []*models.CatBreed `xml:"cat-breed"`
	}

	// Structure the data we want to convert to XML.
	breeds := catBreeds{
		Breeds: allBreeds,
	}

	// Write the XML out.
	_ = t.WriteXML(w, http.StatusOK, breeds)
}

// GetCatBreedByNameXML gets a cat breed from the database, by name, and returns it as XML.
func (app *application) GetCatBreedByNameXML(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "breed")
	var t toolbox.Tools

	// Get all breeds from the database.
	breed, err := app.App.Models.CatBreed.GetBreedByName(name)
	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	// Write the XML out.
	_ = t.WriteXML(w, http.StatusOK, breed)
}

// GetCatBreedByNameJSON gets a cat breed from the database, by name, and returns it as XML.
func (app *application) GetCatBreedByNameJSON(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "breed")
	var t toolbox.Tools

	// Get all breeds from the database.
	breed, err := app.App.Models.CatBreed.GetBreedByName(name)
	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	// Write the XML out.
	_ = t.WriteJSON(w, http.StatusOK, breed)
}
