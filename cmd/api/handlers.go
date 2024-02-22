package main

import (
	"github.com/tsawler/toolbox"
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
