package main

import (
	"github.com/tsawler/toolbox"
	"go-breeders/models"
	"net/http"
)

type Cat struct {
	Breed string `json:"breed"`
}

func (app *application) GetAllCatsJSON(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools
	dogBreeds, err := app.App.Models.CatBreed.All()
	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	_ = t.WriteJSON(w, http.StatusOK, dogBreeds)
}

func (app *application) GetAllCatsXML(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools
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

	breeds := catBreeds{
		Breeds: allBreeds,
	}

	_ = t.WriteXML(w, http.StatusOK, breeds)
}
