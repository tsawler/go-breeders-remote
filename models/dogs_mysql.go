package models

import (
	"context"
	"log"
	"time"
)

func (m *mysqlRepository) GetDogBreedByID(id int) (*DogBreed, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, breed, weight_low_lbs, weight_high_lbs, 
       			cast(((weight_low_lbs + weight_high_lbs) / 2) as unsigned) as average_weight,
				lifespan, coalesce(details, ''),
				coalesce(alternate_names, ''), coalesce(geographic_origin, '') 
				from dog_breeds where id = ?`

	row := m.DB.QueryRowContext(ctx, query, id)
	var dog DogBreed
	err := row.Scan(
		&dog.ID,
		&dog.Breed,
		&dog.WeightLowLbs,
		&dog.WeightHighLbs,
		&dog.AverageWeight,
		&dog.Lifespan,
		&dog.Details,
		&dog.AlternateNames,
		&dog.GeographicOrigin,
	)
	if err != nil {
		log.Println("Error getting breed by id:", err)
		return nil, err
	}
	return &dog, nil
}

func (m *mysqlRepository) RandomDogBreed() (*DogBreed, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, breed, weight_low_lbs, weight_high_lbs,
       			cast(((weight_low_lbs + weight_high_lbs) / 2) as unsigned) as average_weight,
				lifespan, coalesce(details, ''),
				coalesce(alternate_names, ''), coalesce(geographic_origin, '') 
				from dog_breeds order by rand() limit 1`

	row := m.DB.QueryRowContext(ctx, query)
	var dog DogBreed
	err := row.Scan(
		&dog.ID,
		&dog.Breed,
		&dog.WeightLowLbs,
		&dog.WeightHighLbs,
		&dog.AverageWeight,
		&dog.Lifespan,
		&dog.Details,
		&dog.AlternateNames,
		&dog.GeographicOrigin,
	)
	if err != nil {
		log.Println(err)
		return nil, nil
	}

	return &dog, nil
}

func (m *mysqlRepository) AllDogBreeds() ([]*DogBreed, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, breed, weight_low_lbs, weight_high_lbs,
       			cast(((weight_low_lbs + weight_high_lbs) / 2) as unsigned) as average_weight,
				lifespan, coalesce(details, ''),
				coalesce(alternate_names, ''), coalesce(geographic_origin, '') 
				from dog_breeds order by breed`

	var dogs []*DogBreed

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return dogs, err
	}
	defer rows.Close()

	for rows.Next() {
		var d DogBreed
		err := rows.Scan(
			&d.ID,
			&d.Breed,
			&d.WeightLowLbs,
			&d.WeightHighLbs,
			&d.AverageWeight,
			&d.Lifespan,
			&d.Details,
			&d.AlternateNames,
			&d.GeographicOrigin,
		)
		if err != nil {
			return dogs, err
		}
		dogs = append(dogs, &d)
	}
	return dogs, nil
}

// RandomDogOfSize returns a random dog within a specified weight range.
func (m *mysqlRepository) RandomDogOfSize(minWeight, maxWeight int) (*DogBreed, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, breed, weight_low_lbs, weight_high_lbs,
       			cast(((weight_low_lbs + weight_high_lbs) / 2) as unsigned) as average_weight,
				lifespan, coalesce(details, ''),
				coalesce(alternate_names, ''), coalesce(geographic_origin, '') 
				from dog_brees 
				where weight_high_lbs >= ? and weight_high_lbs <= ?
				order by rand() limit 1`

	row := m.DB.QueryRowContext(ctx, query, minWeight, maxWeight)
	var dog DogBreed
	err := row.Scan(
		&dog.ID,
		&dog.Breed,
		&dog.WeightLowLbs,
		&dog.WeightHighLbs,
		&dog.AverageWeight,
		&dog.Lifespan,
		&dog.Details,
		&dog.AlternateNames,
		&dog.GeographicOrigin,
	)
	if err != nil {
		log.Println(err)
		return nil, nil
	}

	return &dog, nil
}
