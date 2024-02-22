package models

func (m *testRepository) GetDogBreedByID(id int) (*DogBreed, error) {
	var dogBreed DogBreed
	return &dogBreed, nil
}

func (m *testRepository) RandomDogBreed() (*DogBreed, error) {
	return nil, nil
}

func (m *testRepository) AllDogBreeds() ([]*DogBreed, error) {
	return nil, nil
}

// RandomDogOfSize returns a random dog within a specified weight range.
func (m *testRepository) RandomDogOfSize(minWeight, maxWeight int) (*DogBreed, error) {
	return nil, nil
}
