package configuration

import (
	"database/sql"
	"go-breeders-remote/models"
	"sync"
)

// Application is the type returned by this package.
type Application struct {
	Models *models.Models
}

// define some variables.
var instance *Application
var once sync.Once
var db *sql.DB

// New is our factory pattern to return a new instance of the app config.
func New(pool *sql.DB) *Application {
	db = pool
	return GetInstance()
}

// GetInstance will always return one, and only one, instance of *Application.
// It uses the sync.Once package to ensure that the initialization of the singleton
// instance only occurs once.
func GetInstance() *Application {
	once.Do(func() {
		instance = &Application{Models: models.New(db)}
	})
	return instance
}
