package models

import (
	"context"
	"log"
	"time"
)

func (m *mysqlRepository) GetCatBreedByID(id int) (*CatBreed, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, breed, weight_low_lbs, weight_high_lbs,
       			cast(((weight_low_lbs + weight_high_lbs) / 2) as unsigned) as average_weight,
				lifespan, coalesce(details, ''),
				coalesce(alternate_names, ''), coalesce(geographic_origin, '') 
				from cat_breeds where id = ?`

	row := m.DB.QueryRowContext(ctx, query, id)
	var cat CatBreed
	err := row.Scan(
		&cat.ID,
		&cat.Breed,
		&cat.WeightLowLbs,
		&cat.WeightHighLbs,
		&cat.AverageWeight,
		&cat.Lifespan,
		&cat.Details,
		&cat.AlternateNames,
		&cat.GeographicOrigin,
	)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &cat, nil
}

func (m *mysqlRepository) GetCatBreedByName(id string) (*CatBreed, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, breed, weight_low_lbs, weight_high_lbs,
       			cast(((weight_low_lbs + weight_high_lbs) / 2) as unsigned) as average_weight,
				lifespan, coalesce(details, ''),
				coalesce(alternate_names, ''), coalesce(geographic_origin, '') 
				from cat_breeds where breed = ?`

	row := m.DB.QueryRowContext(ctx, query, id)
	var cat CatBreed
	err := row.Scan(
		&cat.ID,
		&cat.Breed,
		&cat.WeightLowLbs,
		&cat.WeightHighLbs,
		&cat.AverageWeight,
		&cat.Lifespan,
		&cat.Details,
		&cat.AlternateNames,
		&cat.GeographicOrigin,
	)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &cat, nil
}

func (m *mysqlRepository) RandomCatBreed() (*CatBreed, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, breed, weight_low_lbs, weight_high_lbs,
				lifespan, coalesce(details, ''),
				coalesce(alternate_names, ''), coalesce(geographic_origin, '') 
				from cat_breeds order by rand() limit 1`

	row := m.DB.QueryRowContext(ctx, query)
	var c CatBreed
	err := row.Scan(
		&c.ID,
		&c.Breed,
		&c.WeightLowLbs,
		&c.WeightHighLbs,
		&c.Lifespan,
		&c.Details,
		&c.AlternateNames,
		&c.GeographicOrigin,
	)
	if err != nil {
		log.Println(err)
		return nil, nil
	}

	return &c, nil
}

func (m *mysqlRepository) AllCatBreeds() ([]*CatBreed, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, breed, weight_low_lbs, weight_high_lbs,
       			cast(((weight_low_lbs + weight_high_lbs) / 2) as unsigned) as average_weight,
				lifespan, coalesce(details, ''),
				coalesce(alternate_names, ''), coalesce(geographic_origin, '') 
				from cat_breeds order by breed`

	var cats []*CatBreed

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var c CatBreed
		err := rows.Scan(
			&c.ID,
			&c.Breed,
			&c.WeightLowLbs,
			&c.WeightHighLbs,
			&c.AverageWeight,
			&c.Lifespan,
			&c.Details,
			&c.AlternateNames,
			&c.GeographicOrigin,
		)
		if err != nil {
			return nil, err
		}
		cats = append(cats, &c)
	}
	return cats, nil
}

// RandomCatOfSize returns a random dog within a specified weight range.
func (m *mysqlRepository) RandomCatOfSize(minWeight, maxWeight int) (*CatBreed, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, breed, weight_low_lbs, weight_high_lbs,
				lifespan, coalesce(details, ''),
				coalesce(alternate_names, ''), coalesce(geographic_origin, '') 
				from cat_breeds 
				where weight_high_lbs >= ? and weight_high_lbs <= ?
				order by rand() limit 1`

	row := m.DB.QueryRowContext(ctx, query, minWeight, maxWeight)
	var cat CatBreed
	err := row.Scan(
		&cat.ID,
		&cat.Breed,
		&cat.WeightLowLbs,
		&cat.WeightHighLbs,
		&cat.Lifespan,
		&cat.Details,
		&cat.AlternateNames,
		&cat.GeographicOrigin,
	)
	if err != nil {
		log.Println(err)
		return nil, nil
	}

	return &cat, nil
}
