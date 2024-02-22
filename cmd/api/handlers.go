package main

import (
	"github.com/tsawler/toolbox"
	"go-breeders-remote/models"
	"net/http"
)

// GetAllCatsJSON gets a list of all cat breeds from the database and returns it as JSON.
func (app *application) GetAllCatsJSON(w http.ResponseWriter, r *http.Request) {
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

// GetAllCatsXML gets a list of all cat breeds from the database and returns it as XML.
func (app *application) GetAllCatsXML(w http.ResponseWriter, r *http.Request) {
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
