package main

import (
	"log"
	"net/http"
	"time"
)

const port = ":8000"

type application struct{}

// main is the entry point for our app.
func main() {
	app := application{}

	// create http server
	srv := &http.Server{
		Addr:              port,
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
		WriteTimeout:      30 * time.Second,
	}

	log.Println("*** Starting server on port", port)

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
