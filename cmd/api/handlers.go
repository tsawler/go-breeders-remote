package main

import "net/http"

type Cat struct {
	Breed string `json:"breed"`
}

func (app *application) GetAllCatsJSON(w http.ResponseWriter, r *http.Request) {

}
