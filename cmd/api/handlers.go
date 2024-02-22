package main

import (
	"encoding/xml"
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
	catBreeds, err := app.App.Models.CatBreed.All()
	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	type breeds struct {
		Breeds []*models.CatBreed `xml:"cat-breed"`
	}

	b := breeds{
		Breeds: catBreeds,
	}

	out, err := xml.Marshal(b)
	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/xml")
	w.WriteHeader(http.StatusOK)

	// Add the XML header.
	xmlOut := []byte(xml.Header + string(out))
	_, _ = w.Write(xmlOut)

	//_ = t.WriteXML(w, http.StatusOK, dogBreeds)
}
