package models

import "database/sql"

// Repository is the database repository. Anything that implements
// this interface must implement all the methods included here.
type Repository interface {
	AllDogBreeds() ([]*DogBreed, error)
	RandomDogBreed() (*DogBreed, error)
	GetDogBreedByID(id int) (*DogBreed, error)
	AllCatBreeds() ([]*CatBreed, error)
	RandomCatBreed() (*CatBreed, error)
	GetCatBreedByID(id int) (*CatBreed, error)
	GetCatBreedByName(id string) (*CatBreed, error)
	//RandomDogOfSize(minWeight, maxWeight int) (*models.DogBreed, error)
	//RandomCatOfSize(minWeight, maxWeight int) (*models.CatBreed, error)
}

// mysqlRepository is a simple wrapper for the *sql.DB type. This is
// used to return a MySQL/MariaDB repository.
type mysqlRepository struct {
	DB *sql.DB
}

// newMysqlRepository is a convenience factory method to return a new mysqlRepository.
func newMysqlRepository(conn *sql.DB) Repository {
	return &mysqlRepository{
		DB: conn,
	}
}

// testRepository is a simple wrapper for the *sql.DB type. This is
// used to return a test repository.
type testRepository struct {
	DB *sql.DB
}

// newTestRepository is a convenience factory method to return a new mysqlRepository.
func newTestRepository(conn *sql.DB) Repository {
	return &testRepository{
		DB: nil,
	}
}
