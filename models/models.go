package models

import (
	"database/sql"
	"time"
)

var repo Repository

// Models is the wrapper for all database models.
type Models struct {
	// any models inserted here (and in the New function)
	// are easily accessible throughout the entire application
	DogBreed DogBreed
	CatBreed CatBreed
	Dog      Dog
	Cat      Cat
}

// New initializes the models package for use. We create a DogBreed
// and a CatBreed pre-populated with a repository. This way, we can have
// repository methods for everything, and also methods on CatBreed and DogBreed
// that actually use the repository methods (e.g. DogBreed.Random()). This is another
// example of a simple Factory pattern.
func New(conn *sql.DB) *Models {
	if conn == nil {
		repo = newTestRepository(nil)
	} else {
		repo = newMysqlRepository(conn)
	}

	return &Models{
		DogBreed: DogBreed{},
		CatBreed: CatBreed{},
	}
}

type DogBreed struct {
	ID               int    `json:"id"`
	Breed            string `json:"breed"`
	WeightLowLbs     int    `json:"weight_low_lbs"`
	WeightHighLbs    int    `json:"weight_high_lbs"`
	AverageWeight    int    `json:"average_weight"`
	Lifespan         int    `json:"average_lifespan"`
	Details          string `json:"details"`
	AlternateNames   string `json:"alternate_names"`
	GeographicOrigin string `json:"geographic_origin"`
}

func (d *DogBreed) All() ([]*DogBreed, error) {
	return repo.AllDogBreeds()
}

func (d *DogBreed) Random() (*DogBreed, error) {
	return repo.RandomDogBreed()
}

func (d *DogBreed) Get(id int) (*DogBreed, error) {
	return repo.GetDogBreedByID(id)
}

type CatBreed struct {
	ID               int    `json:"id" xml:"id"`
	Breed            string `json:"breed" xml:"breed"`
	WeightLowLbs     int    `json:"weight_low_lbs" xml:"weight_low_lbs"`
	WeightHighLbs    int    `json:"weight_high_lbs" xml:"weight_high_lbs"`
	AverageWeight    int    `json:"average_weight" xml:"average_weight"`
	Lifespan         int    `json:"average_lifespan" xml:"average_lifespan"`
	Details          string `json:"details" xml:"details"`
	AlternateNames   string `json:"alternate_names" xml:"alternate_names"`
	GeographicOrigin string `json:"geographic_origin" xml:"geographic_origin"`
}

func (c *CatBreed) All() ([]*CatBreed, error) {
	return repo.AllCatBreeds()
}

func (c *CatBreed) Random() (*CatBreed, error) {
	return repo.RandomCatBreed()
}

func (c *CatBreed) Get(id int) (*CatBreed, error) {
	return repo.GetCatBreedByID(id)
}

func (c *CatBreed) GetBreedByName(name string) (*CatBreed, error) {
	return repo.GetCatBreedByName(name)
}

type Dog struct {
	ID             int       `json:"id"`
	DogName        int       `json:"dog_name"`
	BreedID        int       `json:"breed_id"`
	BreederID      int       `json:"breeder_id"`
	Color          string    `json:"color"`
	DateOfBirth    time.Time `json:"date_of_birth"`
	SpayedNeutered int       `json:"spayed_neutered"`
	Description    string    `json:"description"`
	Weight         int       `json:"weight"`
	Breed          DogBreed  `json:"breed"`
	Breeder        Breeder   `json:"breeder"`
}

func (d *Dog) GetBreeder() (*Breeder, error) {
	return &Breeder{}, nil
}

func (d *Dog) DateOfBirthISO() string {
	return d.DateOfBirth.Format("2006-01-02")
}

type Cat struct {
	ID             int       `json:"id"`
	CatName        int       `json:"cat_name"`
	BreedID        int       `json:"breed_id"`
	BreederID      int       `json:"breeder_id"`
	Color          string    `json:"color"`
	DateOfBirth    time.Time `json:"date_of_birth"`
	SpayedNeutered int       `json:"spayed_neutered"`
	Description    string    `json:"description"`
	Weight         int       `json:"weight"`
	Breed          CatBreed  `json:"breed"`
	Breeder        Breeder   `json:"breeder"`
}

func (c *Cat) GetBreeder() (*Breeder, error) {
	return &Breeder{}, nil
}

func (c *Cat) DateOfBirthISO() string {
	return c.DateOfBirth.Format("2006-01-02")
}

type Breeder struct {
	ID          int         `json:"id"`
	BreederName string      `json:"breeder_name"`
	Address     string      `json:"address"`
	City        string      `json:"city"`
	ProvState   string      `json:"prov_state"`
	Country     string      `json:"country"`
	Zip         string      `json:"zip"`
	Phone       string      `json:"phone"`
	Email       string      `json:"email"`
	Active      string      `json:"active"`
	DogBreeds   []*DogBreed `json:"dog_breeds"`
	CatBreeds   []*CatBreed `json:"cat_breeds"`
}

type User struct {
	ID          int
	FirstName   string
	LastName    string
	Email       string
	Password    string
	Active      int
	AccessLevel int
}

// Pet is the type we return from our factory.
type Pet struct {
	Species     string `json:"species"`
	Breed       string `json:"breed"`
	MinWeight   int    `json:"min_weight"`
	MaxWeight   int    `json:"max_weight"`
	Description string `json:"description"`
	LifeSpan    int    `json:"lifespan"`
}
