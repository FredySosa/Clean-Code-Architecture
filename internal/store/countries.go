package store

import (
	"context"

	"github.com/FredySosa/cleanCode/internal/models"
)

type Countries struct {
	countries map[string]models.Country
}

func NewCountryStore(initialData ...map[string]models.Country) Countries {
	data := make(map[string]models.Country, 0)

	if initialData != nil {
		for key, value := range initialData[0] {
			data[key] = value
		}
	}
	return Countries{
		countries: data,
	}
}

func (s Countries) GetCountries(ctx context.Context, data models.QueryData) ([]models.Country, string, error) {
	countries := make([]models.Country, 0)

	for _, country := range s.countries {
		countries = append(countries, country)
	}

	return countries, "0", nil
}
