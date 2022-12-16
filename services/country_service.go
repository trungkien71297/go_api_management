package services

import (
	"go_api_management/domain/countries"
	"go_api_management/utils/errors"
)

var (
	CountryServices countryServicesInterface = &countryServices{}
)

type countryServicesInterface interface {
	GetAllCountries() ([]*countries.Country, *errors.RestError)
}

type countryServices struct{}

func (s *countryServices) GetAllCountries() ([]*countries.Country, *errors.RestError) {
	res, err := countries.GetAllCountries()
	if err != nil {
		return nil, err
	}
	return res, nil
}
