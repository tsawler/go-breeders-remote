package main

import (
	"flag"
	"go-breeders-remote/configuration"
	"log"
	"net/http"
	"time"
)

const port = ":8081"

type application struct {
	App    *configuration.Application // a singleton which is exported, so we can get to it from other modules.
	config appConfig                  // configuration information for the app.
}

// appConfig is a type embedded into the application type. It holds things that no other part of the
// app needs to know about.
type appConfig struct {
	useCache bool
	dsn      string
}

// main is the entry point for our app.
func main() {
	var config appConfig

	// read command line parameters, if any, and set sensible defaults for development
	flag.StringVar(&config.dsn, "dsn", "mariadb:myverysecretpassword@tcp(localhost:3306)/breeders?parseTime=true&tls=false&collation=utf8_unicode_ci&timeout=5s&readTimeout5", "DSN")
	flag.Parse()

	// get database
	db, err := initMySQLDB(config.dsn)
	if err != nil {
		log.Fatal(err)
	}

	app := application{
		App:    configuration.New(db),
		config: config,
	}

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

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
