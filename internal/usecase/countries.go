package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/FredySosa/cleanCode/internal/constants"
	"github.com/FredySosa/cleanCode/internal/models"
)

type (
	Countries struct {
		store CountriesStore
	}
	CountriesStore interface {
		GetCountries(ctx context.Context, data models.QueryData) ([]models.Country, string, error)
	}
)

func NewCountriesUseCase(s CountriesStore) Countries {
	return Countries{
		store: s,
	}
}

func (u Countries) GetCountries(ctx context.Context, limit, offset string) ([]models.Country, string, error) {
	queryData := models.QueryData{
		Offset: offset,
		Limit:  limit,
	}

	countries, newOffset, err := u.store.GetCountries(ctx, queryData)
	if err != nil {
		if errors.Is(err, constants.ErrCountriesNotFound) {
			return []models.Country{}, "", nil
		}
		return nil, "", fmt.Errorf("something went wrong: %w", err)
	}

	return countries, newOffset, nil
}
